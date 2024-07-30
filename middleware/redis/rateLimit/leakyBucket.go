package main

/**
 * @Author huang
 * @Date 2024-07-25
 * @File: leakyBucket.go
 * @Description:
 * 漏桶实现 https://github.com/uber-go/ratelimit
 */

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// redeclare
// var ctx = context.Background()

type LeakyBucket struct {
	client         *redis.Client
	rate           float64 // 每秒处理的请求数
	capacity       int64   // 桶的容量
	lastUpdateTime time.Time
}

func NewLeakyBucket(client *redis.Client, rate float64, capacity int64) *LeakyBucket {
	return &LeakyBucket{
		client:         client,
		rate:           rate,
		capacity:       capacity,
		lastUpdateTime: time.Now(),
	}
}

func (b *LeakyBucket) Allow(key string) bool {
	now := time.Now()
	elapsed := now.Sub(b.lastUpdateTime).Seconds()
	b.lastUpdateTime = now

	// 计算可以处理的请求数
	delta := int64(elapsed * b.rate)

	// 更新桶中的请求数
	currentCount, err := b.client.Get(ctx, key).Int64()
	if err != nil && err != redis.Nil {
		fmt.Println("Error getting key:", err)
		return false
	}

	// 减少已经处理的请求数
	currentCount = max(0, currentCount-delta)

	// 判断桶是否满
	if currentCount < b.capacity {
		// 允许请求并将请求计入桶中
		err = b.client.Set(ctx, key, currentCount+1, 0).Err()
		if err != nil {
			fmt.Println("Error setting key:", err)
			return false
		}
		return true
	} else {
		// 桶满，拒绝请求
		return false
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 初始化 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 初始化漏桶算法
	rate := 1.0           // 每秒处理1个请求
	capacity := int64(10) // 桶的容量为10
	bucket := NewLeakyBucket(client, rate, capacity)

	// 示例请求
	key := "api:leakybucket:20240712:userid"
	for i := 0; i < 15; i++ {
		if bucket.Allow(key) {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond) // 模拟请求间隔
	}
}
