package main

/**
 * @Author huang
 * @Date 2024-07-25
 * @File: tokenBucket.go
 * @Description:
 */

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// redeclare
// var ctx = context.Background()

type TokenBucket struct {
	client         *redis.Client
	rate           float64 // 每秒生成的令牌数
	capacity       int64   // 桶的容量
	lastUpdateTime time.Time
}

func NewTokenBucket(client *redis.Client, rate float64, capacity int64) *TokenBucket {
	return &TokenBucket{
		client:         client,
		rate:           rate,
		capacity:       capacity,
		lastUpdateTime: time.Now(),
	}
}

func (b *TokenBucket) Allow(key string) bool {
	now := time.Now()
	elapsed := now.Sub(b.lastUpdateTime).Seconds()
	b.lastUpdateTime = now

	// 计算可以生成的令牌数
	delta := int64(elapsed * b.rate)

	// 更新桶中的令牌数
	currentTokens, err := b.client.Get(ctx, key).Int64()
	if err != nil && err != redis.Nil {
		fmt.Println("Error getting key:", err)
		return false
	}

	// 增加生成的令牌数，最多不超过桶的容量
	currentTokens = min(b.capacity, currentTokens+delta)

	// 判断桶中是否有足够的令牌处理请求
	if currentTokens > 0 {
		// 允许请求并将令牌计数减1
		err = b.client.Set(ctx, key, currentTokens-1, 0).Err()
		if err != nil {
			fmt.Println("Error setting key:", err)
			return false
		}
		return true
	} else {
		// 没有足够的令牌，拒绝请求
		return false
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 初始化 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 初始化令牌桶算法
	rate := 1.0           // 每秒生成1个令牌
	capacity := int64(10) // 桶的容量为10
	bucket := NewTokenBucket(client, rate, capacity)

	// 示例请求
	key := "api:tokenbucket:20240712:userid"
	for i := 0; i < 15; i++ {
		if bucket.Allow(key) {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
	}
}
