package main

import (
	"fmt"
	"math/rand"
)

/*
*
这个节点实际是链表的头节点而已，用它来表示整个链表(通过它可以到任何一个节点)，但实际上并不是整个链表

1. 为什么使用头节点来代表整个链表？
访问链表的起点：链表是一种线性结构，其元素（节点）的排列有序。链表的操作，无论是添加、删除还是遍历，都需要从链表的起始点开始，这个起始点就是头节点。通过头节点，可以遍历到链表的每一个元素。

简化操作：在很多操作中，尤其是添加和删除节点时，知道头节点的位置可以大大简化逻辑。例如，向链表头部添加一个新节点，只需创建一个新的节点并将其指向原头节点，然后更新头节点为这个新节点即可。

性能优化：对于单向链表，如果有头节点，那么无论链表有多长，访问链表的起始位置的时间复杂度总是O(1)。对于一些特定操作，如在链表头部添加或删除节点，这种设计可以实现高效的执行。为什么使用头节点来代表整个链表？
访问链表的起点：链表是一种线性结构，其元素（节点）的排列有序。链表的操作，无论是添加、删除还是遍历，都需要从链表的起始点开始，这个起始点就是头节点。通过头节点，可以遍历到链表的每一个元素。

简化操作：在很多操作中，尤其是添加和删除节点时，知道头节点的位置可以大大简化逻辑。例如，向链表头部添加一个新节点，只需创建一个新的节点并将其指向原头节点，然后更新头节点为这个新节点即可。

性能优化：对于单向链表，如果有头节点，那么无论链表有多长，访问链表的起始位置的时间复杂度总是O(1)。对于一些特定操作，如在链表头部添加或删除节点，这种设计可以实现高效的执行。

2. 实际上的链表结构
实际上的链表由多个这样的节点通过指针相连组成。每个节点包含两个主要部分：存储的数据（或值）和指向链表中下一个节点的指针（在双向链表中还会有指向前一个节点的指针）。通过这种方式，即使我们只有一个头节点的引用，也能通过遍历操作访问到链表中的每一个节点，从而执行增加、删除、查找等操作。

因此，虽然在代码中我们常常看到只有一个节点的声明来代表整个链表，但实际上整个链表是由这样的多个节点通过“Next”指针串联起来的结构。这种设计既体现了数据结构的逻辑清晰性，也便于在不同的场景下高效地操作链表。
*/

// TODO 查看源码 container/list container/ring container/heap 包

type LinkNode struct {
	Data int8
	Next *LinkNode
}

// 遍历单链表
func traverseLinkedList(head *LinkNode) {
	node := head
	for node != nil {
		fmt.Println("node.Data  ---->  ", node.Data)
		node = node.Next
	}
}

func test1() {
	node1 := new(LinkNode)
	node1.Data = 1

	node2 := new(LinkNode)
	node2.Data = 2
	node1.Next = node2

	node3 := &LinkNode{Data: 3}
	node2.Next = node3

	traverseLinkedList(node1)
}

// //////////////////////////////////////////////////////////////

// Ring 循环链表
// Golang 标准库 container/ring
type Ring struct {
	prev, next *Ring
	Value      int
}

// Ring 初始化
func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

func test2() {
	r := new(Ring)
	r.init()
}

// NewNRings 创建N个节点的循环链表
func NewNRings(n int) *Ring {
	if n <= 0 {
		return nil
	}

	r := &Ring{Value: rand.Int()}
	tmp := r
	for i := 1; i < n; i++ {
		newNode := &Ring{prev: tmp, Value: rand.Int()}
		tmp.next = newNode

		tmp = newNode
	}
	tmp.next = r
	r.prev = tmp
	fmt.Println("r----", r, r.prev, r.Value, r.next)
	r1 := r.next
	fmt.Println("r1----", r1, r1.prev, r1.Value, r1.next)
	r2 := r1.next
	fmt.Println("r2----", r2, r2.prev, r2.Value, r2.next)
	/**
	result below from console:
	r---- &{0x1400000c078 0x1400000c060 2246265063424630037} &{0x1400000c060 0x1400000c048 8784267196909056036} 2246265063424630037 &{0x1400000c048 0x1400000c078 764836491650020610}
	r1---- &{0x1400000c048 0x1400000c078 764836491650020610} &{0x1400000c078 0x1400000c060 2246265063424630037} 764836491650020610 &{0x1400000c060 0x1400000c048 8784267196909056036}
	r2---- &{0x1400000c060 0x1400000c048 8784267196909056036} &{0x1400000c048 0x1400000c078 764836491650020610} 8784267196909056036 &{0x1400000c078 0x1400000c060 2246265063424630037}
	*/
	return r
}

// Next 获取下一个节点
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// Move 获取第 n 个节点
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// Link 添加节点 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

// Unlink 删除节点后面的 n 个节点
func (r *Ring) Unlink(n int) *Ring {
	if n < 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// //////////////////////////////////////////////////////////////////////

// 数组和链表是两个不同的概念。一个是编程语言提供的基本数据类型，表示一个连续的内存空间，可通过一个索引访问数据。另一个是我们定义的数据结构，通过一个数据节点，可以定位到另一个数据节点，不要求连续的内存空间。
//
// 数组的优点是占用空间小，查询快，直接使用索引就可以获取数据元素，缺点是移动和删除数据元素要大量移动空间。
//
// 链表的优点是移动和删除数据元素速度快，只要把相关的数据元素重新链接起来，但缺点是占用空间大，查找需要遍历。
//
// 很多其他的数据结构都由数组和链表配合实现的。

func arr() {
	array := [5]int64{}
	fmt.Println(array)
	array[0] = 8
	array[1] = 9
	array[2] = 7
	fmt.Println(array)
	fmt.Println(array[2])
}

func ArrayLink() {
	type Value struct {
		Data      string
		NextIndex int64
	}

	var array [5]Value          // 五个节点的数组
	array[0] = Value{"I", 3}    // 下一个节点的下标为3
	array[1] = Value{"Army", 4} // 下一个节点的下标为4
	array[2] = Value{"You", 1}  // 下一个节点的下标为1
	array[3] = Value{"Love", 2} // 下一个节点的下标为2
	array[4] = Value{"!", -1}   // -1表示没有下一个节点
	node := array[0]
	for {
		fmt.Println(node.Data)
		if node.NextIndex == -1 {
			break
		}
		node = array[node.NextIndex]
	}

}

func main() {
	test1()
	// test2()
	// NewNRings(3)
	// linkNewTest()
}
