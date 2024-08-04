package main

import (
	"fmt"
	"github.com/kardianos/osext"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	// getCurrentPath()
	// getCurrentPath1()

	// getExecutingFIlePath1()

	// getExecutingFIlePath2()
}
