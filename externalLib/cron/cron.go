package main

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: cron.go
 * @Description:
 */

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 创建一个新的 cron 实例
	c := cron.New()

	// 添加一个定时任务 (每 3 s运行一次)
	_, err := c.AddFunc("@every 3s", func() {
		fmt.Println("任务运行于 ---------------------------->", time.Now())
		// 在这里执行你的任务
	})
	if err != nil {
		fmt.Println("添加定时任务失败:", err)
		return
	}

	// 启动定时任务调度器 (不会阻塞主线程)
	c.Start()

	// 确保在程序退出时停止定时任务
	defer c.Stop()

	select {}
}
