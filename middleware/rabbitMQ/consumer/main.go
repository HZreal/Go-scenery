package main

import (
	"log"
	"time"

	"goScenery/middleware/rabbitMQ/internal/rabbitmqdemo"
)

func main() {
	cfg, err := rabbitmqdemo.ConfigFromEnv()
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

	if err := ch.Qos(1, 0, false); err != nil {
		log.Fatalf("set qos failed: %v", err)
	}

	deliveries, err := ch.Consume(
		queue.Name,
		"go-rabbitmq-demo-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("consume failed: %v", err)
	}

	log.Printf("listening for messages from %s through %s", queue.Name, cfg.URL)

	for delivery := range deliveries {
		log.Printf("received: %s", delivery.Body)
		time.Sleep(300 * time.Millisecond)
		if err := delivery.Ack(false); err != nil {
			log.Printf("ack failed: %v", err)
		}
	}
}
