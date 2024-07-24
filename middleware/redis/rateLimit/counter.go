package main

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: counter.go
 * @Description:
 */

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

var ctx = context.Background()

// 初始化 Redis 客户端
var client = redis.NewClient(&redis.Options{
	Addr: "localhost:56378",
	DB:   15,
})

func rateLimit(key string, limit int) bool {
	// 检查键是否存在
	count, err := client.Get(ctx, key).Int64()
	if errors.Is(err, redis.Nil) {
		// 键不存在，进行初始化
		err = client.Set(ctx, key, 1, 60*time.Second).Err()
		if err != nil {
			fmt.Println("Error setting key:", err)
			return false
		}
		return true
	} else if err != nil {
		// 处理其他错误
		fmt.Println("Error getting key:", err)
		return false
	}

	// 键存在，进行自增操作
	count, err = client.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println("Error incrementing key:", err)
		return false
	}

	// 检查计数器值是否超过阈值
	if int(count) > limit {
		// 超过阈值，拒绝请求
		return false
	}

	// 未超过阈值，允许请求
	return true
}

func main() {

	// 定义限流阈值
	limit := 100

	// 示例请求
	key := "api:limit:20240712:userid"
	if !rateLimit(key, limit) {
		fmt.Println("Request denied")
	}

	fmt.Println("Request allowed")
}
