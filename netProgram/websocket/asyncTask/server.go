package main

/**
 * @Author elastic·H
 * @Date 2024-10-24
 * @File: server.go
 * @Description:
 */

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 定义任务结构
type Task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

// 用于存储任务的 map 和锁
var (
	taskStore = make(map[string]*Task)
	mu        sync.Mutex
)

// 生成唯一任务ID
func generateTaskID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// 处理 WebSocket 连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		// 读取客户端发送的任务配置
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read message error:", err)
			break
		}

		log.Printf("Received task config: %s", message)

		// 创建并存储异步任务
		taskID := generateTaskID()
		task := &Task{
			ID:     taskID,
			Status: "created",
			Result: "",
		}

		mu.Lock()
		taskStore[taskID] = task
		mu.Unlock()

		// 向客户端返回初始任务信息
		err = conn.WriteJSON(task)
		if err != nil {
			log.Println("Write JSON error:", err)
			break
		}

		// 模拟异步任务执行
		go executeTask(taskID, conn)
	}
}

// 模拟执行异步任务
func executeTask(taskID string, conn *websocket.Conn) {
	time.Sleep(10 * time.Second) // 模拟任务执行时间

	mu.Lock()
	task, exists := taskStore[taskID]
	if exists {
		task.Status = "completed"
		task.Result = "Task result for ID: " + taskID
	}
	mu.Unlock()

	// 推送任务完成信息
	if exists {
		err := conn.WriteJSON(task)
		if err != nil {
			log.Println("Push task result error:", err)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	log.Println("Server started at :52222")
	err := http.ListenAndServe(":52222", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
