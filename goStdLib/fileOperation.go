package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var filename1 = "goStdLib/123.txt"
var filename2 = "goStdLib/456.txt"

func useIOUtilToRW() {

	// 读文件
	file, _ := ioutil.ReadFile(filename1)
	fmt.Println(file) // []byte

	strFile := string(file)
	fmt.Println(strFile)

	// 写文件
	str := "hello\nworld!\t1234\ncat"
	_ = ioutil.WriteFile(filename2, []byte(str), 0644)

}

func useOSToRead() {

	// 打开文件   指定模式、权限打开文件
	file, err := os.OpenFile(filename2, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 打开文件(实际是调用指定模式和权限的 OpenFile)
	file1, err := os.Open(filename1)
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	// 读取文件
	b1 := make([]byte, 12)  // 定义长度为12的切片，来存储file1中读取的字节数据
	n1, _ := file1.Read(b1) // 读取file1中长度为len(b)的字节
	fmt.Printf("读取的字节数: %d\n读取的字节保存在切片b中: %s\n", n1, string(b1))

	newOffset, _ := file1.Seek(0, 0)
	fmt.Println("文件指针移动后的位置--------", newOffset)

	// io.ReadAtLeast读取打开的file1
	b2 := make([]byte, 5)
	n2, _ := io.ReadAtLeast(file1, b2, 3)
	fmt.Println(n2, string(b2))

	// bufio.NewReader通过file1创建reader缓冲区，将文件中的内容预加载到缓存中，方便快速读取出来。当然，文件的内容非常多的时候，它是一部分一部分加载到缓冲区的，并不会将所有内容一次全部加载完
	reader := bufio.NewReader(file)
	//reader.Read([]byte)
	//_, _ =reader.ReadBytes('\n')
	//b3, _ := reader.Peek(4)
	//fmt.Println(string(b3))
	for {
		//读取一行数据
		buf, err := reader.ReadBytes('\n') // 参数delim为分隔符，每次读到遇到分隔符停止，若在找到分隔符之前遇到错误(通常是io.EOF)，它会返回在错误和错误本身之前读到的数据
		if err != nil && err == io.EOF {
			fmt.Println("读完了所有数据")
			fmt.Println(string(buf))
			return
		} else if err != nil {
			fmt.Println("ReaderBytes 读取出错:", err)
			return
		}
		fmt.Println(string(buf))
	}

}

func useOSToWrite() {
	// 创建文件(实际是调用指定模式和权限的 OpenFile)
	file, err := os.Create("goStdLib/789.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	//n1, _ := file.Write([]byte("hello123"))
	//n2, _ := file.WriteString("\nworld\n456")    // 内部调用的还是 *File.Write
	//fmt.Println("两次写入的字节长度-----", n1, n2)

	// bufio.NewWriter通过file1创建writer，再写入
	writer := bufio.NewWriter(file)
	n3, _ := writer.Write([]byte("hello123"))
	n4, _ := writer.WriteString("\nworld\n456")
	fmt.Println("两次写入的字节长度-----", n3, n4)
	_ = writer.Flush() // 写入的数据还在缓冲区，需要刷回硬盘

}

func main() {
	//useIOUtilToRW()
	useOSToRead()
	//useOSToWrite()
}
