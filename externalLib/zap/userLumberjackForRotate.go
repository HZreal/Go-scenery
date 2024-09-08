package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)

/**
 * @Author elastic·H
 * @Date 2024-09-08
 * @File: userLumberjackForRotate.go
 * @Description: 使用 Lumberjack 进行日志 rotate
 */

var sugarLogger2 *zap.SugaredLogger

func init() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger2 = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/zapLumberjackLog.log", // 日志文件位置
		MaxSize:    1,                           // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxBackups: 5,                           // 保留旧文件的最大个数
		MaxAge:     30,                          // 保留旧文件的最大天数
		Compress:   false,                       // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	sugarLogger2.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger2.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger2.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	defer sugarLogger2.Sync()
	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}
