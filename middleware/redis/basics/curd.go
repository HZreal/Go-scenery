package main

/**
 * @Author huang
 * @Date 2023-06-15
 * @File: curd.go
 * @Description:
 */

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func init() {
	//
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:56378",
		Password: "", // no password set
		DB:       15,
	})
}

func keyNotExist() {
	// key2 does not exist
	val2, err := client.Get(ctx, "keyNotExist").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("keyNotExist", val2)
	}
}

func rwString() {
	// 操作 String
	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)
}

func rwList() {
	// 操作 List
	err := client.RPush(ctx, "list", "value1", "value2", "value3").Err()
	if err != nil {
		panic(err)
	}

	listVal, err := client.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("list:", listVal)
}

func rwHash() {
	// 操作 Hash
	err := client.HSet(ctx, "hash", "field1", "value1", "field2", "value2").Err()
	if err != nil {
		panic(err)
	}

	hashVal, err := client.HGetAll(ctx, "hash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("hash:", hashVal)
}

func rwSet() {
	// 操作 Set
	err := client.SAdd(ctx, "set", "member1", "member2", "member3").Err()
	if err != nil {
		panic(err)
	}

	setVal, err := client.SMembers(ctx, "set").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("set:", setVal)
}

func rwZSet() {
	// 操作 ZSet
	err := client.ZAdd(ctx, "zset", redis.Z{Score: 1, Member: "member1"}, redis.Z{Score: 2, Member: "member2"}).Err()
	if err != nil {
		panic(err)
	}

	zSetVal, err := client.ZRangeWithScores(ctx, "zset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZSet:", zSetVal)
}
func main() {
	keyNotExist()
	rwString()
	rwList()
	rwHash()
	rwSet()
	rwZSet()
}
