package main

/**
 * @Author elastic·H
 * @Date 2024-08-04
 * @File: notServerCase.go
 * @Description: 非 Server 情况的分析
 */

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
)

var currentDir string

// 当前文件所在目录
func getCurrentDir() {
	// 当前文件所在目录
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("can not get current file info"))
	}
	fmt.Println("filePath ---->  ", filePath)
	currentDir = filepath.Dir(filePath)
}

func init() {
	getCurrentDir()
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func case1() {

	// CPU profiling
	cpuFile, err := os.Create(filepath.Join(currentDir, "notServerCpu.prof"))
	if err != nil {
		fmt.Println("could not create CPU profile: ", err)
		return
	}
	defer cpuFile.Close()

	// 开始 CPU 分析
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Println("could not start CPU profile: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	// 运行一些计算密集型的代码
	for i := 0; i < 30; i++ {
		fibonacci(35)
	}

	// Memory profiling
	memFile, err := os.Create(filepath.Join(currentDir, "notServerMemory.prof"))
	if err != nil {
		fmt.Println("could not create memory profile: ", err)
		return
	}
	defer memFile.Close()

	// runtime.GC() // 在获取堆信息前先进行一次GC

	// 开始内存分析
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Println("could not write memory profile: ", err)
		return
	}

	fmt.Println("CPU and Memory profiles created")
}
func main() {
	case1()
}

// 执行 go run xxx.go 后，在代码目录下，可以看到生成了 cpu.prof 和 memory.prof 文件。
// 然后使用 go tool pprof 工具进行性能分析
// 		有 2 中分析方式：
// 			1. 终端命令交互方式：执行 go tool pprof /path/to/memory.prof
//			2. WEB UI 方式：go tool pprof -http=:8888 /path/to/memory.prof
