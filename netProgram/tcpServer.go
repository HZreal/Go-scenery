package main

import (
	"bufio"
	"fmt"
	"goBasics/netProgram/proto"
	"io"
	"net"
)

//TCP服务端程序的处理流程：
//		1.监听端口
//		2.接收客户端请求建立链接
//		3.创建goroutine处理链接。

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte             // 定义缓冲，作为读取数据的存储
		n, err := reader.Read(buf[:]) // 读取数据，放入buf
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte("服务器返还给客户端的数据")) // 将数据发送给客户端
	}

}

func server1() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listener.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

// tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发会出现粘包
// “粘包”可发生在发送端也可发生在接收端：
//    	1.由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
//    	2.接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。
// 解决： 出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作
// 封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。

// 我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度
// 创建proto包，及Decode和Encode函数处理数据

func process2(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func server2() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process2(conn)
	}
}
func main() {
	//server1()
	server2()
}
