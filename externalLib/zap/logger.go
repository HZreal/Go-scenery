package main

/**
 * @Author elastic·H
 * @Date 2024-09-08
 * @File: logger.go
 * @Description:
 */

import (
	"go.uber.org/zap"
	"time"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
	// logger, _ = zap.NewProduction()
}

func useLogger() {
	logger.Info("User logged in",
		zap.String("username", "Alice"),
		zap.Time("timestamp", time.Now()),
	)

	//
	// logger.Warn("Failed to login",
	// 	zap.String("username", "Alice"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
}

func main() {
	useLogger()
}
