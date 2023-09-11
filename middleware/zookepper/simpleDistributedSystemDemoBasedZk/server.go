package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"net"
	"os"
	"time"
)

// 对于一个分布式服务，最基本的一项功能就是服务的注册和发现，而利用zk的EPHEMERAL节点则可以很方便的实现该功能
// EPHEMERAL节点正如其名，是临时性的，其生命周期是和客户端会话绑定的，当会话连接断开时，节点也会被删除

// 基于zk实现一个简单的分布式server如下：

// server启动时，创建zk连接，并在/go_servers/节点下创建一个新节点，节点名为"ip:port"，完成服务注册
// server结束时，由于连接断开，创建的节点会被删除，这样client就不会连到该节点

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetConnect() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	checkError(err)
	return
}

func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func handleCient(conn net.Conn, port string) {
	fmt.Println("-------------", port)
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(port + ": " + daytime)) // 返回数据给client
}

// 在zk中/go_server/下注册三个节点，根据addr创建3个TCP服务端连接，处理客户端的请求
func starServer(addr string) {
	// 1. 服务注册(将启动的三个TCP服务注册到zk中保存)
	// 连接zk
	conn, err := GetConnect()
	checkError(err)
	defer conn.Close()
	// zk节点注册(实际就是在/go_server/下创建一个node为addr)
	err = RegistServer(conn, addr)
	checkError(err)

	// 2. 启动tcp server监听addr(ip端口)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	fmt.Println(tcpAddr)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn2, err := listener.Accept() // 挂载，等待tcp客户端连接
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		go handleCient(conn2, addr) // 获得连接，进开启一个协程，异步处理此连接，继续挂载等待其他连接
	}

}

func main() {
	go starServer("127.0.0.1:8897")
	go starServer("127.0.0.1:8898")
	go starServer("127.0.0.1:8899")

	// keep main-thread alive
	a := make(chan bool, 1)
	<-a
}
