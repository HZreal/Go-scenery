package main

/**
 * @Author elastic·H
 * @Date 2024-07-31
 * @File: goToolTrace.go
 * @Description:
 */

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	// 创建trace文件
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	fmt.Println("Hello World")
}

// 通过 go tool trace 工具打开
// go tool trace trace.out 通过浏览器打开并点击 view trace 查看
