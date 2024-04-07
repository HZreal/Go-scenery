package main

import "fmt"

var _list = []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}

func selectionSort(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ {

		// 令此轮的第 1 个元素为最小值
		min := list[i]
		minIndex := i

		// 从此轮的第 2 个元素开始遍历，查找比 min 更小的值
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		// 只要此轮遍历存在比 min 更小的值，将最小值交换到此轮的第一个元素上
		if minIndex != i {
			list[minIndex], list[i] = list[i], list[minIndex]
		}
	}
}
func ttest1() {
	_list1 := _list
	selectionSort(_list1)
	fmt.Println(_list1)
}

func main() {
	fmt.Println("原始数组 ==>  ", _list)
	ttest1()
}
