package main

import (
	"fmt"
	"sync"
)

// 栈：先进后出，先进队的数据最后才出来。在英文的意思里，stack 可以作为一叠的意思，这个排列是垂直的，你将一张纸放在另外一张纸上面，先放的纸肯定是最后才会被拿走，因为上面有一张纸挡住了它。
// 队列：先进先出，先进队的数据先出来。在英文的意思里，queue 和现实世界的排队意思一样，这个排列是水平的，先排先得。

// 可以用数据结构：链表（可连续或不连续的将数据与数据关联起来的结构），或 数组（连续的内存空间，按索引取值） 来实现 栈（stack） 和 队列 (queue)。
// 		数组实现：能快速随机访问存储的元素，通过下标 index 访问，支持随机访问，查询速度快，但存在元素在数组空间中大量移动的操作，增删效率低。
// 		链表实现：只支持顺序访问，在某些遍历操作中查询速度慢，但增删元素快。

// ArrayStack 数组栈，后进先出
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

// Push 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)

	// 栈中元素数量+1
	stack.size++
}

// Pop 出栈
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		// 空栈
		panic("empty stack!")
	}

	// 栈顶元素
	top := stack.array[stack.size-1]

	// 法1: 通过切片的方式收缩，去掉 array 尾部元素
	// 可能占用空间越来越大
	// stack.array = stack.array[0 : stack.size-1]

	// 法2: 新建数组，保存去除元素后的数组
	// 空间占用不会越来越大，但可能移动元素次数过多
	newArr := make([]string, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArr[i] = stack.array[i]
	}
	stack.array = newArr
	// var newArr2 []string
	// for i := 0; i < stack.size-1; i++ {
	//	newArr2 = append(newArr2, stack.array[i])
	// }
	// stack.array = newArr2

	stack.size--

	return top
}

// Peek 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}

// GetSize 栈大小
func (stack *ArrayStack) GetSize() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.GetSize())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.GetSize())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
