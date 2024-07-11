package main

/**
 * @Author huang
 * @Date 2024-07-11
 * @File: subscribe.go
 * @Description:
 */

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:56378",
		DB:   15, // 使用 15 数据库
	})

	pubsub := rdb.Subscribe(ctx, "test_channel")
	defer pubsub.Close()

	// 等待订阅确认
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatalf("Could not subscribe to channel: %v", err)
	}

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Printf("Received message from channel: %s\n", msg.Payload)
		// 在这里处理消息
	}
}
