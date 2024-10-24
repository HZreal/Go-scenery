package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"sync"
	"time"
)

// Constants
const (
	// Service types
	ServiceTypeChat         = "chat"
	ServiceTypeNotification = "notification"
	ServiceTypeGame         = "game"

	// Client types
	ClientTypeWeb    = "web"
	ClientTypeMobile = "mobile"
	ClientTypeUE     = "ue"

	// Redis keys
	RedisKeyConnPrefix = "ws:conn:"

	// Timeouts
	WriteTimeout = 10 * time.Second
	ReadTimeout  = 60 * time.Second
	PingPeriod   = 54 * time.Second
)

// Message represents a WebSocket message
type Message struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

// Connection represents a WebSocket connection
type Connection struct {
	ID           string
	UserID       string
	ClientType   string
	ServiceType  string
	Conn         *websocket.Conn
	Send         chan []byte
	Hub          *Hub
	LastPingTime time.Time
	mu           sync.Mutex
}

// Hub manages all active connections
type Hub struct {
	connections    map[string]*Connection
	register       chan *Connection
	unregister     chan *Connection
	broadcast      chan []byte
	redis          *redis.Client
	messageHandler MessageHandler
	mu             sync.RWMutex
}

// MessageHandler interface for processing different message types
type MessageHandler interface {
	Handle(conn *Connection, msg Message) error
}

// Config represents server configuration
type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

// NewHub creates a new Hub instance
func NewHub(redis *redis.Client, handler MessageHandler) *Hub {
	return &Hub{
		connections:    make(map[string]*Connection),
		register:       make(chan *Connection),
		unregister:     make(chan *Connection),
		broadcast:      make(chan []byte),
		redis:          redis,
		messageHandler: handler,
	}
}

// Run starts the hub's main loop
func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.mu.Lock()
			h.connections[conn.ID] = conn
			h.mu.Unlock()

			// Store connection info in Redis
			err := h.storeConnectionInfo(conn)
			if err != nil {
				log.Printf("Failed to store connection info: %v", err)
			}

		case conn := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.connections[conn.ID]; ok {
				delete(h.connections, conn.ID)
				close(conn.Send)
			}
			h.mu.Unlock()

			// Remove connection info from Redis
			err := h.removeConnectionInfo(conn)
			if err != nil {
				log.Printf("Failed to remove connection info: %v", err)
			}

		case message := <-h.broadcast:
			h.mu.RLock()
			for _, conn := range h.connections {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(h.connections, conn.ID)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// storeConnectionInfo stores connection information in Redis
func (h *Hub) storeConnectionInfo(conn *Connection) error {
	ctx := context.Background()
	key := RedisKeyConnPrefix + conn.ID

	connInfo := map[string]interface{}{
		"user_id":      conn.UserID,
		"client_type":  conn.ClientType,
		"service_type": conn.ServiceType,
		"created_at":   time.Now().Unix(),
	}

	return h.redis.HMSet(ctx, key, connInfo).Err()
}

// removeConnectionInfo removes connection information from Redis
func (h *Hub) removeConnectionInfo(conn *Connection) error {
	ctx := context.Background()
	key := RedisKeyConnPrefix + conn.ID
	return h.redis.Del(ctx, key).Err()
}

// Upgrader specifies parameters for upgrading an HTTP connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Implement your origin check logic here
		return true
	},
}

// AuthMiddleware handles authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate token here
		// ...

		next.ServeHTTP(w, r)
	})
}

// handleWebSocket handles WebSocket connections
func (h *Hub) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	clientType := r.URL.Query().Get("client_type")
	serviceType := r.URL.Query().Get("service_type")
	userID := r.URL.Query().Get("user_id") // Should be extracted from token in production

	if !isValidClientType(clientType) || !isValidServiceType(serviceType) {
		http.Error(w, "Invalid client_type or service_type", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	connection := &Connection{
		ID:          fmt.Sprintf("%s-%s-%s", userID, clientType, serviceType),
		UserID:      userID,
		ClientType:  clientType,
		ServiceType: serviceType,
		Conn:        conn,
		Send:        make(chan []byte, 256),
		Hub:         h,
	}

	h.register <- connection

	go connection.writePump()
	go connection.readPump()
}

// writePump pumps messages from the hub to the WebSocket connection
func (c *Connection) writePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(WriteTimeout))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(WriteTimeout))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// readPump pumps messages from the WebSocket connection to the hub
func (c *Connection) readPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512)
	c.Conn.SetReadDeadline(time.Now().Add(ReadTimeout))
	c.Conn.SetPongHandler(func(string) error {
		c.mu.Lock()
		c.LastPingTime = time.Now()
		c.mu.Unlock()
		c.Conn.SetReadDeadline(time.Now().Add(ReadTimeout))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		if err := c.Hub.messageHandler.Handle(c, msg); err != nil {
			log.Printf("Failed to handle message: %v", err)
		}
	}
}

// Helper functions
func isValidClientType(clientType string) bool {
	return clientType == ClientTypeWeb || clientType == ClientTypeMobile || clientType == ClientTypeUE
}

func isValidServiceType(serviceType string) bool {
	return serviceType == ServiceTypeChat || serviceType == ServiceTypeNotification || serviceType == ServiceTypeGame
}

// DefaultMessageHandler implements MessageHandler interface
type DefaultMessageHandler struct{}

func (h *DefaultMessageHandler) Handle(conn *Connection, msg Message) error {
	switch msg.Type {
	case "chat":
		return h.handleChatMessage(conn, msg)
	case "notification":
		return h.handleNotificationMessage(conn, msg)
	case "game":
		return h.handleGameMessage(conn, msg)
	default:
		return errors.New("unknown message type")
	}
}

func (h *DefaultMessageHandler) handleChatMessage(conn *Connection, msg Message) error {
	// Implement chat message handling
	return nil
}

func (h *DefaultMessageHandler) handleNotificationMessage(conn *Connection, msg Message) error {
	// Implement notification message handling
	return nil
}

func (h *DefaultMessageHandler) handleGameMessage(conn *Connection, msg Message) error {
	// Implement game message handling
	return nil
}

func main() {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:56379",
		Password: "", // no password set
		DB:       10, // use default DB
	})

	// Create hub with default message handler
	hub := NewHub(rdb, &DefaultMessageHandler{})
	go hub.Run()

	// Set up HTTP server with WebSocket endpoint
	http.Handle("/ws", AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hub.handleWebSocket(w, r)
	})))

	// Start server
	log.Fatal(http.ListenAndServe(":52222", nil))
}

// ws://your-domain:8080/ws?client_type={client_type}&service_type={service_type}&user_id={user_id}
