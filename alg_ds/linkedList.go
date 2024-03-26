package main

import (
	"fmt"
	"math/rand"
)

type LinkNode struct {
	Data int8
	Next *LinkNode
}

func test1() {
	node1 := new(LinkNode)
	node1.Data = 1

	node2 := new(LinkNode)
	node2.Data = 2
	node1.Next = node2

	node3 := LinkNode{Data: 3}
	node2.Next = &node3

	nowNode := node1
	for {
		if nowNode != nil {
			fmt.Printf("%T  ", nowNode)
			fmt.Println("nowNode.Data = ", nowNode.Data)
			nowNode = nowNode.Next
			continue
		}
		break
	}
}

/*
*
定义链表
*/
type Ring struct {
	prev, next *Ring
	value      int
}

/*
*
初始化 Ring
*/
func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

func test2() {
	r := new(Ring)
	r.init()
}

// 创建N个节点的循环链表
func newNRings(n int) *Ring {
	if n <= 0 {
		return nil
	}

	r := &Ring{value: rand.Int()}
	tmp := r
	for i := 1; i < n; i++ {
		newNode := &Ring{prev: tmp, value: rand.Int()}
		tmp.next = newNode

		tmp = newNode
	}
	tmp.next = r
	r.prev = tmp
	fmt.Println("r----", r, r.prev, r.value, r.next)
	r1 := r.next
	fmt.Println("r1----", r1, r1.prev, r1.value, r1.next)
	r2 := r1.next
	fmt.Println("r2----", r2, r2.prev, r2.value, r2.next)
	/**
	result below from console:
	r---- &{0x1400000c078 0x1400000c060 2246265063424630037} &{0x1400000c060 0x1400000c048 8784267196909056036} 2246265063424630037 &{0x1400000c048 0x1400000c078 764836491650020610}
	r1---- &{0x1400000c048 0x1400000c078 764836491650020610} &{0x1400000c078 0x1400000c060 2246265063424630037} 764836491650020610 &{0x1400000c060 0x1400000c048 8784267196909056036}
	r2---- &{0x1400000c060 0x1400000c048 8784267196909056036} &{0x1400000c048 0x1400000c078 764836491650020610} 8784267196909056036 &{0x1400000c078 0x1400000c060 2246265063424630037}
	*/
	return r
}
func main() {
	//test1()
	//test2()
	newNRings(3)
}
