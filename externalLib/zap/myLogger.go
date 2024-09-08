package main

/**
 * @Author elastic·H
 * @Date 2024-09-08
 * @File: myLogger.go
 * @Description:
 */

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var mySugarLogger *zap.SugaredLogger
var myLogger *zap.Logger

func init() {
	// 定义 writeSyncer
	// file, _ := os.Create("logs/zapLogs.log")
	file, _ := os.OpenFile("logs/zapLogs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writeSyncer := zapcore.AddSync(file)

	// 定义 encoder
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	// 定义 level
	level := zap.DebugLevel

	// 基于上述 3 者创建 core
	core := zapcore.NewCore(encoder, writeSyncer, level)
	// 基于 core 创建自定义 logger
	logger := zap.New(core)
	mySugarLogger = logger.Sugar()
}

func userMyLogger() {
	mySugarLogger.Infof("User %s logged in at %v", "Alice", time.Now())
}

func main() {
	userMyLogger()
}
