package main

/**
 * @Author nico
 * @Date 2026-01-21
 * @File: concurrenceWrite.go
 * @Description:
 */

import (
	"fmt"
	"sync"
)

func writeConcurrently() {
	var wg sync.WaitGroup
	resultCh := make(chan string, 1000) // 带缓冲更好

	workers := 10
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			// 模拟工作...
			for j := 0; j < 100; j++ {
				resultCh <- fmt.Sprintf("worker-%d-item-%d", id, j)
			}
		}(i)
	}

	// 关闭 channel 的正确时机：所有生产者都完成后
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 收集结果
	var results []string
	for v := range resultCh {
		results = append(results, v)
	}

	fmt.Printf("总共收集到 %v 条数据\n", results)
}

func main() {
	writeConcurrently()
}
