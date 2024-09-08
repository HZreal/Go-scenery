package main

/**
 * @Author elastic·H
 * @Date 2024-09-08
 * @File: sugarLogger.go
 * @Description:
 */

import (
	"go.uber.org/zap"
	"time"
)

var sugarLogger *zap.SugaredLogger

func init() {
	logger1, _ := zap.NewDevelopment()
	// logger1, _ := zap.NewProduction()

	sugarLogger = logger1.Sugar()
}

func useSugarLogger() {
	sugarLogger.Infof("User %s logged in at %v", "Alice", time.Now())

	// Warnw 方法使用键值对记录结构化日志，增加可读性
	sugarLogger.Warnw("Failed to login",
		"username", "Alice",
		"attempt", 3,
		"backoff", time.Second,
	)
}

func main() {
	useSugarLogger()
}
