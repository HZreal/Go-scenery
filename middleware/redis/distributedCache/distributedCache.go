package main

/**
 * @Author huang
 * @Date 2024-07-25
 * @File: distributedCache.go
 * @Description:
 */

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

func main() {
	clusterAddrs := []string{
		"192.168.1.100:7000",
		"192.168.1.101:7001",
		"192.168.1.102:7002",
		"192.168.1.103:7003",
		"192.168.1.104:7004",
		"192.168.1.105:7005",
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        clusterAddrs,
		Password:     "", // Redis 密码（如果有）
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("无法连接到 Redis 集群：%v\n", err)
		return
	}
	fmt.Printf("连接成功：%s\n", pong)

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Printf("写操作失败：%v\n", err)
		return
	}
	fmt.Println("写操作成功")

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		fmt.Printf("读操作失败：%v\n", err)
		return
	}
	fmt.Printf("读取的值：%s\n", val)
}
