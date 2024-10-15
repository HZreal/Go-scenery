package main

/**
 * @Author elastic·H
 * @Date 2024-10-15
 * @File: funcImplInterface.go
 * @Description:
 */

import (
	"fmt"
)

type Greet interface {
	SayHello(name string)
}

type greetFunc func(string)

// 实现接口方法
func (f greetFunc) SayHello(name string) {
	f(name) // 调用函数本身
}

func main() {
	// 定义一个符合 greetFunc 类型的函数
	// 匿名函数转换为 greetFunc 类型
	myGreet := greetFunc(func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	})

	// 将函数作为 Greet 接口的实现使用
	var greeter Greet = myGreet
	greeter.SayHello("Alice")
}
