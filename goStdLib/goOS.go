package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 获取当前的执行路径
// /private/var/folders/s8/k10vrs290_l9hr6kr4lrmckh0000gq/T
func getCurrentPath() {
	pathStr, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(pathStr)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	fmt.Println(ret)
}

// 获取当前的执行路径
// C:\Users\Vic\AppData\Local\Temp\
func getCurrentPath1() {
	str, err := exec.LookPath(os.Args[0])
	checkErr(err)
	fmt.Println("str", str)
	i := strings.LastIndex(str, "\\") // 仅windows
	path := string(str[0 : i+1])
	fmt.Println(path)
}

// 获取当前文件的详细路径
func CurrentFile() {
	_, file, _, ok := runtime.Caller(1) // 不要放在main函数里调用
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	fmt.Println(file)
}

func main() {
	// fmt.Println(os.Getwd())

	// getCurrentPath()
	// getCurrentPath1()
	// CurrentFile()

	_, file, _, _ := runtime.Caller(1) // 在main函数中调用，返回执行时路径
	fmt.Println(file)                  // 输出/Users/hz/Desktop/hz/go1.17.8/go1.17.8/src/runtime/proc.go

}
