package main

import (
	"fmt"
	"strconv"
)

// strconv包主要实现基本数据类型与其字符串表示的转换

// Atoi() 该函数用于将字符串类型的整数转换为int类型
func atoiFunc() {
	number, _ := strconv.Atoi("123")
	fmt.Println(number)
}

// Itoa() 该函数用于将int类型数据数据转换为对应的字符串表现形式
func itoaFunc() {
	str := strconv.Itoa(123)
	fmt.Println(str)
}

// TODO Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()
func parseSerialsFunc() {

}

// TODO Format系列函数实现了将给定类型数据格式化为string类型数据的功能
func formatSerialsFunc() {

}

func main() {
	atoiFunc()
	// itoaFunc()
	// parseSerialsFunc()
	// formatSerialsFunc()
}
