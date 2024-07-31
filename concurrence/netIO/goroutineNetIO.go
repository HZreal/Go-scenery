package main

/**
 * @Author elastic·H
 * @Date 2024-07-31
 * @File: goroutineNetIO.go
 * @Description:
 */

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sync() {
	// 发送HTTP GET请求
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	// 打印响应体
	fmt.Println(string(body))
}

// fetchURL 通过协程发送HTTP GET请求，并通过通道通知主程序
func fetchURL(ch chan<- string) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading body: %v", err)
		return
	}

	ch <- string(body)
}

func async() {
	ch := make(chan string)

	// 启动协程发送网络请求
	go fetchURL(ch)

	// 等待协程完成并接收通知
	result := <-ch
	fmt.Println(result)
}
func main() {
	// sync()
	async()
}
