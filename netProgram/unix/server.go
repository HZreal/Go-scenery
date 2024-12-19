package main

/**
 * @Author nico
 * @Date 2024-12-19
 * @File: server.go
 * @Description:
 */

import (
	"log"
	"net"
	"os"
)

func main() {
	SocketPath := "assets/go_unix_socket"

	// 删除之前的套接字文件（如果存在）
	if err := os.RemoveAll(SocketPath); err != nil {
		log.Println("Failed to remove existing socket:", err)
		return
	}

	// 创建监听器
	listener, err := net.Listen("unix", SocketPath)
	if err != nil {
		log.Println("Failed to listen on socket:", err)
		return
	}
	defer listener.Close()

	log.Println("Server is listening on", SocketPath)

	// 等待客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		defer conn.Close()

		log.Println("Client connected")

		// 读取客户端发送的数据
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Failed to read from client:", err)
			return
		}

		// 打印客户端发送的数据
		log.Printf("Received: %s\n", string(buf[:n]))

		// 向客户端发送响应
		response := "Hello from server!"
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Println("Failed to send response to client:", err)
			return
		}
	}
}
