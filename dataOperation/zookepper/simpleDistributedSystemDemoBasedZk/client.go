package main

import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"

	// "github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"math/rand"
	"net"
	"time"
)

// 先从zk获取go_servers节点下所有子节点，这样就拿到了所有注册的server 从server列表中选中一个节点（这里只是随机选取，实际服务一般会提供多种策略），创建连接进行通信
// 这里为了演示，我们每次client连接server，获取server发送的时间后就断开。

func checkError1(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetConnect1() (conn *zk.Conn, err error) {
	zkList := []string{"localhost:2181"}
	conn, _, err = zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}

// 连接zk，获取随机某一个(实际的分布式系统会有算法选择某个addr)
func getServerHost() (host string, err error) {
	conn, err := GetConnect1()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()
	serverList, err := GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty \n")
		return
	}

	// 随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}

// 根据此分布式系统选中的addr创建TCP客户端连接，并发送数据
func startClient() {
	// 1. 获取要请求的cluster node server addr
	serverHost, err := getServerHost()
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}
	fmt.Println("connect host: " + serverHost)

	// 2. 通过addr创建TCP客户端连接
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	checkError1(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError1(err)
	defer conn.Close()
	// 发送数据
	_, err = conn.Write([]byte("timestamp"))
	checkError1(err)
	// 获取结果
	result, err := ioutil.ReadAll(conn)
	checkError1(err)
	fmt.Println("result -------->", string(result))

	return
}

func main() {
	for i := 0; i < 100; i++ { // 模拟大量客户端请求
		startClient()
		time.Sleep(1 * time.Second)
	}
}
