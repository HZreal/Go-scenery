package main

/**
 * @Author nico
 * @Date 2025-03-17
 * @File: copier.go
 * @Description:
 */

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type A struct {
	Field1 string
	Field2 int
	Field3 bool
}

type B struct {
	Field1 string
	Field2 int
	Field3 bool
	Field4 float64
	Field5 string
}

func copyStruct() {
	a := A{Field1: "hello", Field2: 42, Field3: true}
	b := B{Field4: 3.14, Field5: "world"}

	// 将 A 的字段复制到 B 中
	err := copier.Copy(&b, &a)
	if err != nil {
		fmt.Println("复制失败:", err)
		return
	}

	fmt.Println(b) // 输出: {hello 42 true 3.14 world}
}

func copySlice() {
	aa := []A{
		{Field1: "hello", Field2: 42, Field3: true},
		{Field1: "world", Field2: 100, Field3: false},
	}

	var bb []B

	// 将 aa 复制到 bb 中
	err := copier.Copy(&bb, &aa)
	if err != nil {
		fmt.Println("复制失败:", err)
		return
	}

	// 设置 B 的额外字段
	for i := range bb {
		bb[i].Field4 = 3.14
		bb[i].Field5 = "extra"
	}

	fmt.Println(bb)
	// 输出: [{hello 42 true 3.14 extra} {world 100 false 3.14 extra}]
}

func main() {
	copyStruct()
	copySlice()
}
