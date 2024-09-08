package main

/**
 * @Author elastic·H
 * @Date 2024-09-08
 * @File: goLog.go
 * @Description:
 */

import (
	"log"
	"os"
)

// 	•	log.Print(v ...interface{})：输出日志信息，不带时间戳。
//	•	log.Println(v ...interface{})：输出日志信息，自动添加空格并带时间戳。
//	•	log.Printf(format string, v ...interface{})：格式化输出日志信息。
//	•	log.Fatal(v ...interface{})：输出日志信息并调用 os.Exit(1) 退出程序。
//	•	log.Fatalf(format string, v ...interface{})：格式化输出日志并退出程序。
//	•	log.Panic(v ...interface{})：输出日志信息并引发 panic。
//	•	log.Panicf(format string, v ...interface{})：格式化输出日志并引发 panic。

// log 包是并发安全的，多协程同时写日志时不需要额外的锁机制。因为 log 内部已经处理了并发的同步问题

// basic
func basic() {
	// 一般输出
	log.Print("Log with Print")
	log.Println("Log with Println")
	log.Printf("Log with %s", "Printf")

	// log.Fatal 会调用 os.Exit(1) 终止运行，即后面的代码不会允许
	log.Fatal("Log with Fatal")
	// log.Fatalf("Log with %s", "Fatalf")

	// log.Panic 会抛出异常
	// log.Panic("Log with Panic")
	// log.Panicf("Log with %s", "Panicf")
}

// logWithSetting
func logWithSetting() {
	// 通过 SetFlags 和 SetPrefix 定制日志格式。
	log.SetPrefix("INFO: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a log message.")
}

// logToFile
func logToFile() {
	file, err := os.OpenFile("logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("log message to file")
}

func myLogger() {
	file, err := os.OpenFile("custom.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建自定义 Logger
	logger := log.New(file, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("This is a custom log message.")
}

func main() {
	// basic()
	// logWithSetting()
	logToFile()
}
