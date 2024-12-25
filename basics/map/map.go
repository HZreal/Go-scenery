package main

/**
 * @Author HZreal
 * @Date 2024-12-17
 * @File: map.go
 * @Description:
 */

import (
	"fmt"
	"sync"
)

// ---------------------------------------------------------------------------------------
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

// ---------------------------------------------------------------------------------------
// slice 和 map 都是引用类型，它们包含对底层数据结构的指针。因此，当你将它们作为参数传递、返回或存储时，需要特别小心，以避免不必要的数据共享或修改

type Stats struct {
	mu sync.Mutex

	counters map[string]int
}

// Snapshot1 返回 map 的引用，外部可以直接修改
// ！！！这种方式不可取
func (s *Stats) Snapshot1() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.counters
}

// Snapshot2
// 拷贝返回，避免外部的修改直接影响到原数据
// ！！！采用这种方式
func (s *Stats) Snapshot2() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}

// ---------------------------------------------------------------------------------------

func main() {
	demo1()
	demo2()
}
