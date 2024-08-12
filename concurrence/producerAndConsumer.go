package main

/**
 * @Author elastic·H
 * @Date 2024-08-13
 * @File: producerAndConsumer.go
 * @Description:
 */

import (
	"fmt"
	"sync"
)

type task struct {
	id int
}

func produce(ch chan<- task, count int) {
	for i := 0; i < count; i++ {
		ch <- task{id: i}
	}
	close(ch)
}

func createConsumer(ch <-chan task, num int, wg *sync.WaitGroup) {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for msg := range ch {
				fmt.Println(id, msg.id)
			}
		}(i)
	}
}

func oneProduceMultiConsumer() {
	ch := make(chan task, 10)

	wg = sync.WaitGroup{}
	createConsumer(ch, 5, &wg)
	produce(ch, 50)
	wg.Wait()
}
