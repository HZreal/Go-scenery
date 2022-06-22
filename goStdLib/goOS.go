package main

import (
	"errors"
	"fmt"
	"github.com/kardianos/osext"
	"log"
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

func getWd() {
	path, err := os.Getwd()
	checkErr(err)
	fmt.Println(path)
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
	fmt.Println("filePath ---->  ", file)
	dir := filepath.Dir(file)
	fmt.Println("dirPath ---->  ", dir)
}

func getExecutingFIlePath1() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// 可执行文件的路径
	fmt.Println(ex)

	//	获取执行文件所在目录
	exPath := filepath.Dir(ex)
	fmt.Println("可执行文件所在目录路径 :" + exPath)
}

func getExecutingFIlePath2() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("编译后二进制文件的执行目录 ----  ", folderPath)
}

func main() {
	// getWd()

	// getCurrentPath()
	// getCurrentPath1()

	CurrentFile()

	// _, file, _, _ := runtime.Caller(1) // 在main函数中调用，返回执行时路径
	// fmt.Println(file)                  // 输出/Users/hz/Desktop/hz/goSDK/go1.17.11/src/runtime/proc.go

	// getExecutingFIlePath1()

	// getExecutingFIlePath2()
}
