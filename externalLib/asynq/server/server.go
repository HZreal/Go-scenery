package main

/**
 * @Author elastic·H
 * @Date 2024-11-06
 * @File: server.go
 * @Description:
 */

import (
	"github.com/hibiken/asynq"
	"goScenery/externalLib/asynq/config"
	"goScenery/externalLib/asynq/tasks"
	"log"
)

func main() {
	// 初始化 Asynq 服务端
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.RedisAddr, DB: config.DB},
		asynq.Config{
			Concurrency: config.Concurrency, // 并发任务数
		},
	)

	// 路由任务处理器
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, tasks.HandleEmailDeliveryTask)
	// mux.Handle(tasks.TypeEmailDelivery, tasks.NewEmailDeliveryProcessor())

	// 启动服务
	if err := server.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
