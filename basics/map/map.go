package main

/**
 * @Author HZreal
 * @Date 2024-12-17
 * @File: map.go
 * @Description:
 */

import (
	"fmt"
)

// 在进行Assignments(分配、赋值)时，左侧操作数必须是：可寻址的、映射索引表达式 或 空白标识符

type Temp struct {
	id int32
}

// 错误示范
func demo1() {
	m := make(map[int32]Temp)

	// 给 map 赋值
	m[1] = Temp{1}
	fmt.Println(m[1])

	// 给map里的结构体的成员赋值
	// m[1].id = 2 // 报错: Cannot assign to m[1].id
	// fmt.Println(m[1].id)

	// 当使用映射（map）的索引表达式时，如果映射的值类型是一个值类型（非指针），那么你不能直接对索引出来的值取地址。
	// vPtr := &m[1] // 报错: Cannot take the address of 'm[1]'
	v := m[1]
	vPtr := &v
	vPtr.id = 2
}

// 正确方式
func demo2() {
	m := make(map[int32]*Temp)

	// 给 map 赋值
	m[1] = &Temp{1}
	fmt.Println(m[1]) // {1}

	// 给 map 里的指针结构体的成员赋值
	m[1].id = 2
	fmt.Println(m[1].id) // 2

}

func main() {
	demo1()
	demo2()
}
