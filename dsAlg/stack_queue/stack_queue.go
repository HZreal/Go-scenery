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
// 时间复杂度为 O(1)
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
	// 操作的时间复杂度为：O(1)
	// stack.array = stack.array[0 : stack.size-1]

	// 法2: 新建数组，保存去除元素后的数组
	// 空间占用不会越来越大，但可能移动元素次数过多
	// 时间复杂度为：O(n)
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

func test1() {
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

// //////////////////////////////////////////////////////////////

// LinkNode 链栈节点
type LinkNode struct {
	Value interface{}
	Next  *LinkNode
}

// LinkStack 链栈/*
type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// Size 获取栈大小
func (stack *LinkStack) Size() int {
	return stack.size
}

// IsEmpty 判空
func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

// Peek 获取栈顶元素
func (stack *LinkStack) Peek() interface{} {
	if stack.size == 0 {
		panic("Empty stack")
	}

	return stack.root.Value
}

// Push 入栈
// 时间复杂度为：O(1)
func (stack *LinkStack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.root == nil {
		// 为空栈，新增
		stack.root = new(LinkNode)
		stack.root.Value = value
	} else {
		// 非空栈，将新增的节点置为头节点
		preTop := stack.root

		// 新节点
		newTop := new(LinkNode)
		newTop.Value = value

		// 原来的链表链接到新元素后面
		newTop.Next = preTop

		// 将新节点放在头部
		stack.root = newTop
	}

	// 栈中元素数量+1
	stack.size++
}

// Pop 出栈
// 时间复杂度为：O(1)
func (stack *LinkStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 空栈无法 Pop
	if stack.root == nil {
		panic("Empty stack")
	}

	// 取栈顶节点值
	topNode := stack.root.Value

	// 将栈顶元素的后继节点作为栈的栈顶节点
	stack.root = stack.root.Next

	//
	stack.size--

	return topNode
}

func test2() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Size())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
}

// //////////////////////////////////////////////////////////////
// 单调栈

// //////////////////////////////////////////////////////////////

// ArrayQueue 数组队列
//
//	head  <<<<<<< ArrayQueue <<<<<<< tail
type ArrayQueue struct {
	array []interface{}
	size  int
	lock  sync.Mutex
}

// Add 入队，从尾部添加
// 时间复杂度为：O(n)
func (queue *ArrayQueue) Add(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, v)

	queue.size++
}

// Remove 出队，从头部移除
// 时间复杂度是：O(n)
func (queue *ArrayQueue) Remove() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("Empty Queue")
	}

	head := queue.array[0]

	// 切片，但空间不会被释放
	// queue.array = queue.array[1:queue.size]

	// 创建新数组
	newArr := make([]interface{}, queue.size-1, queue.size-1)
	for i := 0; i < queue.size-1; i++ {
		newArr[i] = queue.array[i+1]
	}
	queue.array = newArr

	//
	queue.size--

	return head
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *ArrayQueue) Size() int {
	return queue.size
}

func test3() {

}

// //////////////////////////////////////////////////////////////

// LinkQueue 链式队列
//
//	head  <<<<<<< LinkQueue <<<<<<< tail
type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// Add 入队，从尾部添加
// 需要遍历链表，时间复杂度为：O(n)
func (queue *LinkQueue) Add(v interface{}) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 创建新增的节点
	newNode := new(LinkNode)
	newNode.Value = v

	if queue.size == 0 {
		// 空队列，直接新建
		queue.root = newNode
	} else {
		// 非空队列，在尾部添加

		// 先遍历找到尾节点
		nowNode := queue.root
		if nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		// 将新节点链接到尾部节点
		nowNode.Next = newNode
	}

	//
	queue.size++
}

// Remove 出队，从头部移除
// 链表第一个节点出队，时间复杂度为：O(1)
func (queue *LinkQueue) Remove() interface{} {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("Empty queue")
	}

	headNode := queue.root
	v := headNode.Value

	queue.root = headNode.Next
	queue.size--

	return v
}

func (queue *LinkQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func test4() {

}

func main() {
	test1()
	test2()
	test3()
	test4()
}
