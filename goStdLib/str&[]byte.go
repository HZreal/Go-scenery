package main

import "fmt"

// byte 类型  实际是uint8类型别名

// string  字符串是一系列8位字节的集合，通常但不一定代表UTF-8编码的文本。字符串可以为空，但不能为nil。而且字符串的值是不能改变的
// string 底层定义的是一个结构体(一个指针指向数组头，一个表示数组长度)
//string的指针指向的内容是不可以更改的，所以每更改一次字符串，就得重新分配一次内存，之前分配空间的还得由gc回收，这是导致string操作低效的根本原因

// []byte  字节数组  底层定义的是一个切片数组

//string和[]byte，底层都是数组
//[]byte比string灵活，拼接性能也更高

//如何取舍？
//既然string就是一系列字节，而[]byte也可以表达一系列字节，那么实际运用中应当如何取舍？
//		string可以直接比较，而[]byte不可以，所以[]byte不可以当map的key值。
//		因为无法修改string中的某个字符，需要粒度小到操作一个字符时，用[]byte。
//		string值不可为nil，所以如果你想要通过返回nil表达额外的含义，就用[]byte。
//		[]byte切片这么灵活，想要用切片的特性就用[]byte。
//		需要大量字符串处理的时候用[]byte，性能好很多。

func StrToByteArr() {
	str := "hello world!"
	byteArr := []byte(str)
	fmt.Printf("type is %T\nvalue is %v\n", byteArr, byteArr)
}

func ByteArrToStr() {
	byteArr := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', '!'}
	fmt.Println(byteArr)
	str := string(byteArr)
	fmt.Println(str)
}

func main() {
	s := "A1" // 分配存储"A1"的内存空间，s结构体里的str指针指向这快内存
	s = "A2"  // 重新给"A2"的分配内存空间，s结构体里的str指针指向这快内存
	fmt.Println(s)

	//其实[]byte和string的差别是更改变量的时候array的内容可以被更改
	ss := []byte{1} // 分配存储1数组的内存空间，ss结构体的array指针指向这个数组。
	ss = []byte{2}  // 将array的内容改为2
	fmt.Println(ss)

	// 相互转化
	StrToByteArr()
	ByteArrToStr()
}
