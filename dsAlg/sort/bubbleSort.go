package main

import "fmt"

var list = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

// 原始经典冒泡
// 冒泡到尾部
func bubbleSort1(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ { // n-1 轮遍历
		for j := 0; j < n-i-1; j++ {

			// 将最大值冒泡到尾部
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}

			// 将最小值冒泡到尾部
			// if list[j] < list[j+1] {
			// 	list[j], list[j+1] = list[j+1], list[j]
			// }
		}
	}
}

// 冒泡到头部
func bubbleSort2(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {

			// 将最小值冒泡到头部
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}

			// 将最大值冒泡到头部
			// if list[j] > list[j-1] {
			// 	list[j], list[j-1] = list[j-1], list[j]
			// }
		}
	}

}

// BubbleSort_ 改进的冒泡
func BubbleSort_(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ {
		didSwap := false

		for j := 0; j < n-i-1; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}

		// 如果在一轮中没有交换过，那么已经排好序了，直接返回
		if !didSwap {
			return
		}
	}
}

func BubbleSort(list []int) {
	n := len(list)

	// 进行 n-1 轮迭代
	for i := n - 1; i > 0; i-- {
		// 在一轮中有没有交换过
		didSwap := false

		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}

		// 如果在一轮中没有交换过，那么已经排好序了，直接返回
		if !didSwap {
			return
		}
	}
}

func test1() {
	list1 := list
	bubbleSort1(list1)
	fmt.Println(list1)
	list2 := list
	bubbleSort2(list2)
	fmt.Println(list2)
}

func test2() {
	list3 := list
	BubbleSort_(list3)
	fmt.Println(list3)

	list4 := list
	BubbleSort(list4)
	fmt.Println(list4)

}

func main() {
	fmt.Println("原始数组 ==>  ", list)
	test1()
	fmt.Println("--------------------------------")
	test2()
}
