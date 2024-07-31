package main

/**
 * @Author elastic·H
 * @Date 2024-08-01
 * @File: createProcess.go
 * @Description:
 */

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func t1() {
	// 创建一个执行外部命令的命令对象
	cmd := exec.Command("ls", "-l")

	// 设置环境变量
	cmd.Env = append(os.Environ(), "FOO=bar")

	// 设置标准输入输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 运行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func t2() {
	// 创建一个执行外部命令的命令对象
	cmd := exec.Command("sleep", "5")

	// 异步启动命令
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 在等待进程完成的同时做其他事情
	fmt.Println("Process started, waiting for it to finish...")

	// 等待进程完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Process finished with error:", err)
	} else {
		fmt.Println("Process finished successfully")
	}
}

func t3() {
	// 创建一个执行外部命令的命令对象
	cmd := exec.Command("grep", "hello")

	// 获取输入输出管道
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 向标准输入写入数据
	stdin.Write([]byte("hello world\nhello go\n"))
	stdin.Close()

	// 读取标准输出的数据
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 等待命令完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func t4() {
	// 要执行的命令和参数
	execPath := "/bin/echo"
	args := []string{"echo", "Hello, World!"}

	// 设置环境变量
	env := os.Environ()

	// 设置进程属性
	attr := &os.ProcAttr{
		Dir: "",  // 工作目录
		Env: env, // 环境变量
		Files: []*os.File{
			os.Stdin,  // 标准输入
			os.Stdout, // 标准输出
			os.Stderr, // 标准错误
		},
		Sys: &syscall.SysProcAttr{
			// 在这里可以设置更多的系统属性
		},
	}

	// 启动新进程
	process, err := os.StartProcess(execPath, args, attr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 等待进程退出
	state, err := process.Wait()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出进程退出状态
	fmt.Println("my subProcess exited with status:", state)
}

func main() {
	// t1()
	// t2()
	// t3()
	t4()
}
