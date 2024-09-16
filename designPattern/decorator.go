package main

/**
 * @Author elastic·H
 * @Date 2024-09-16
 * @File: decorator.go
 * @Description: 装饰器模式
 */

import (
	"fmt"
)

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Operation before func")
		f(s)
		fmt.Println("Operation after func")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func main() {
	hello := decorator(Hello)
	hello("Hello, World!")
}
