package main

/**
 * @Author huang
 * @Date 2024-07-11
 * @File: distributeLock.go
 * @Description:
 */

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

var ctx = context.Background()

// Redis 客户端
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:56378", // Redis 地址
	DB:   15,                // 使用 15 数据库
})

// 获取锁
func acquireLock(key string, expiration time.Duration) (string, bool) {
	value := uuid.New().String()
	success, err := redisClient.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		log.Printf("Failed to acquire lock: %v", err)
		return "", false
	}
	if success {
		return value, true
	}
	return "", false
}

// 释放锁
func releaseLock(key string, value string) bool {
	// 释放锁 Lua 脚本
	var releaseLockScript = `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return 0
	end
	`
	result, err := redisClient.Eval(ctx, releaseLockScript, []string{key}, value).Result()
	if err != nil {
		log.Printf("Failed to release lock: %v", err)
		return false
	}
	return result.(int64) == 1
}

func test1() {
	key := "myLock"
	expiration := 10 * time.Second

	value, success := acquireLock(key, expiration)
	if !success {
		log.Println("Failed to acquire lock")
	}

	//
	log.Println("Lock acquired, value:", value)

	// 处理你的业务逻辑
}

// /////////////////////////////////////////////////////////////////////////////////////////////

/*
*
模拟两个协程抢锁
*/
func test2() {
	var wg sync.WaitGroup

	key := "myLock"
	expiration := 10 * time.Second

	// 协程 1: 获取锁并释放锁
	wg.Add(1)
	go func() {
		defer wg.Done()

		value, success := acquireLock(key, expiration)
		if success {
			fmt.Println("Goroutine 1: Lock acquired, value:", value)
			// 模拟业务处理
			time.Sleep(5 * time.Second)
			// 释放锁
			if releaseLock(key, value) {
				fmt.Println("Goroutine 1: Lock released")
			} else {
				fmt.Println("Goroutine 1: Failed to release lock")
			}
		} else {
			fmt.Println("Goroutine 1: Failed to acquire lock")
		}
	}()

	// 协程 2: 尝试获取锁并释放锁
	wg.Add(1)
	go func() {
		defer wg.Done()

		// 等待 1 秒，以确保第一个协程已经获取到锁
		time.Sleep(1 * time.Second)

		value, success := acquireLock(key, expiration)
		if success {
			fmt.Println("Goroutine 2: Lock acquired, value:", value)
			// 模拟业务处理
			time.Sleep(5 * time.Second)
			// 释放锁
			if releaseLock(key, value) {
				fmt.Println("Goroutine 2: Lock released")
			} else {
				fmt.Println("Goroutine 2: Failed to release lock")
			}
		} else {
			fmt.Println("Goroutine 2: Failed to acquire lock")
		}
	}()

	// 等待所有协程完成
	wg.Wait()
}

// /////////////////////////////////////////////////////////////////////////////////////////////

// 续租锁
func renewLock(key string, value string, expiration time.Duration) bool {
	renewScript := `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("PEXPIRE", KEYS[1], ARGV[2])
	else
		return 0
	end
	`
	result, err := redisClient.Eval(ctx, renewScript, []string{key}, value, expiration.Milliseconds()).Result()
	if err != nil {
		log.Printf("Failed to renew lock: %v", err)
		return false
	}
	return result.(int64) == 1
}

// 模拟业务处理逻辑
func businessLogic(key string, value string, wg *sync.WaitGroup) {
	defer wg.Done()
	expiration := 10 * time.Second

	if !renewLock(key, value, expiration) {
		fmt.Println("Failed to renew lock, exiting")
		return
	}

	// 模拟业务处理时间
	workTime := 15 * time.Second
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		time.Sleep(workTime)
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("Business logic completed")
			return
		case <-ticker.C:
			if !renewLock(key, value, expiration) {
				fmt.Println("Failed to renew lock during processing")
				return
			}
		}
	}
}

/*
*
模拟多个协程抢锁并执行业务逻辑，且执行业务逻辑期间续租锁
*/
func test3() {
	var wg sync.WaitGroup
	key := "myLock"

	// 模拟 5 个抢锁
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			expiration := 10 * time.Second
			value, success := acquireLock(key, expiration)
			if success {
				fmt.Printf("Goroutine %d: Lock acquired, value: %s\n", id, value)

				wg.Add(1)
				businessLogic(key, value, &wg)

				if releaseLock(key, value) {
					fmt.Printf("Goroutine %d: Lock released\n", id)
				} else {
					fmt.Printf("Goroutine %d: Failed to release lock\n", id)
				}
			} else {
				fmt.Printf("Goroutine %d: Failed to acquire lock\n", id)
			}
		}(i)
	}

	wg.Wait()
}

// /////////////////////////////////////////////////////////////////////////////////////////////

// 获取锁
func acquireLock2(key, value string, expiration time.Duration) bool {
	result, err := redisClient.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		log.Printf("Failed to acquire lock: %v", err)
		return false
	}
	return result
}

// 续租锁
func renewLock2(key, value string, expiration time.Duration, stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			currentValue, err := redisClient.Get(ctx, key).Result()
			if err != nil {
				log.Printf("Failed to get lock: %v", err)
				return
			}
			if currentValue == value {
				_, err := redisClient.PExpire(ctx, key, expiration).Result()
				if err != nil {
					log.Printf("Failed to renew lock: %v", err)
					return
				}
				log.Printf("Lock renewed for key: %s", key)
			} else {
				log.Printf("Lock value mismatch for key: %s", key)
				return
			}
		case <-stopChan:
			return
		}
	}
}

// 模拟业务逻辑
func businessLogic2(key, value string, wg *sync.WaitGroup) {
	var lockExpire = 10 * time.Second
	defer wg.Done()

	// 获取锁
	if !acquireLock2(key, value, lockExpire) {
		log.Printf("Failed to acquire lock for key: %s", key)
		return
	}
	log.Printf("Lock acquired for key: %s", key)

	// 续租锁
	stopChan := make(chan struct{})
	var renewWg sync.WaitGroup
	renewWg.Add(1)
	go renewLock2(key, value, lockExpire, stopChan, &renewWg)

	// 模拟业务处理
	time.Sleep(20 * time.Second) // 假设业务逻辑处理需要20秒

	// 业务逻辑完成，停止续租
	close(stopChan)
	renewWg.Wait()

	// 释放锁
	if !releaseLock(key, value) {
		log.Printf("Failed to release lock for key: %s", key)
		return
	}
	log.Printf("Lock released for key: %s", key)
}

func test4() {
	var wg sync.WaitGroup
	wg.Add(1)
	go businessLogic2("myLock", "unique-value", &wg)
	wg.Wait()
}

// /////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
