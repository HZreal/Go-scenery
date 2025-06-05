package main

/**
 * @Author nico
 * @Date 2025-06-05
 * @File: splitTest.go
 * @Description:
 */

import (
	"fmt"
	"strings"
)

func t() {
	// 基本用法
	fmt.Println(strings.Split("a,b,c", ",")) // [a b c]

	// 分隔符不存在
	fmt.Println(strings.Split("abc", ",")) // [abc]

	// 空字符串
	fmt.Println(strings.Split("", ",")) // [""]

	// 分隔符为空字符串
	fmt.Println(strings.Split("abc", "")) // [a b c]

	// 字符串和分隔符都为空
	fmt.Println(strings.Split("", "")) // []

	// 连续分隔符
	fmt.Println(strings.Split("a,,b,c", ",")) // [a  b c]

	// 以分隔符开头
	fmt.Println(strings.Split(",a,b", ",")) // [ a b]

	// 以分隔符结尾
	fmt.Println(strings.Split("a,b,", ",")) // [a b ]

	// 只包含分隔符
	fmt.Println(strings.Split(",,,", ",")) // [   ]

	// Unicode字符
	fmt.Println(strings.Split("你-好-世界", "-")) // [你 好 世界]
}

func main() {
	t()
}
