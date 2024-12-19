package main

/**
 * @Author nico
 * @Date 2024-12-19
 * @File: client.go
 * @Description:
 */

import (
	"log"
	"net"
)

func main() {
	SocketPath := "assets/go_unix_socket"

	// 连接到服务端
	conn, err := net.Dial("unix", SocketPath)
	if err != nil {
		log.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	// 发送数据到服务端
	message := "Hello from client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("Failed to send message to server:", err)
		return
	}

	// 读取服务端的响应
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Failed to read response from server:", err)
		return
	}

	// 打印服务端响应
	log.Printf("Server response: %s\n", string(buf[:n]))
}
