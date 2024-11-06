package main

import (
	"goScenery/externalLib/asynq/config"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"goScenery/externalLib/asynq/tasks"
)

var client *asynq.Client

func init() {
	client = asynq.NewClient(asynq.RedisClientOpt{Addr: config.RedisAddr, DB: config.DB})
}

// 任务入队 立即执行
func test1(task *asynq.Task) {
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

// 任务入队 延迟执行
func test2(task *asynq.Task) {
	info, err := client.Enqueue(task, asynq.ProcessIn(10*time.Second))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

// 其他可选项
// Options include MaxRetry, Queue, Timeout, Deadline, Unique etc.
func test3(task *asynq.Task) {
	info, err := client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func main() {
	defer client.Close()

	task1, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task1: %v", err)
	}

	// 立即执行任务
	// test1(task1)

	// 延迟执行任务
	test2(task1)

	// task2, err := tasks.NewImageResizeTask("https://example.com/myassets/image.jpg")
	// if err != nil {
	// 	log.Fatalf("could not create task2: %v", err)
	// }
	// test3(task2)

}
