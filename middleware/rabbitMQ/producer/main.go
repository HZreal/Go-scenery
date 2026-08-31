package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"goScenery/middleware/rabbitMQ/internal/rabbitmqdemo"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg, err := rabbitmqdemo.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	messageCount, err := rabbitmqdemo.ProducerMessageCountFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := rabbitmqdemo.Dial(cfg.URL)
	if err != nil {
		log.Fatalf("connect rabbitmq through haproxy failed: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("open channel failed: %v", err)
	}
	defer ch.Close()

	queue, err := rabbitmqdemo.DeclareDemoQueue(ch, cfg.Queue)
	if err != nil {
		log.Fatalf("declare queue failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 1; i <= messageCount; i++ {
		body := fmt.Sprintf("rabbitmq cluster demo message %02d at %s", i, time.Now().Format(time.RFC3339))
		err = ch.PublishWithContext(
			ctx,
			"",
			queue.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			},
		)
		if err != nil {
			log.Fatalf("publish message %d failed: %v", i, err)
		}
		log.Printf("published: %s", body)
	}
}
