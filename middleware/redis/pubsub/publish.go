package main

/**
 * @Author huang
 * @Date 2024-07-11
 * @File: publish.go
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

	err := rdb.Publish(ctx, "test_channel", "Hello, Redis!").Err()
	if err != nil {
		log.Fatalf("Could not publish message: %v", err)
	}

	fmt.Println("Message published to channel")
}
