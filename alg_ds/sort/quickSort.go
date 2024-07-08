package main

/**
 * @Author huang
 * @Date 2024-06-08
 * @File: quickSort.go
 * @Description:
 * 参考: https://www.hello-algo.com/chapter_sorting/quick_sort/
 */

import (
	"fmt"
)

// partition函数用于分区，并返回基准元素的位置
func partition(arr []int, left, right int) int {
	pivot := arr[left] // 选择第一个元素作为基准元素
	i := left + 1
	j := right

	for i <= j {
		// 向右移动 i 指针，直到找到一个大于或等于 pivot 的元素
		for i <= j && arr[i] <= pivot {
			i++
		}
		// 向左移动 j 指针，直到找到一个小于或等于 pivot 的元素
		for i <= j && arr[j] >= pivot {
			j--
		}
		if i < j {
			// 交换 i 和right 指针指向的元素
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	// 将基准元素放到正确的位置
	arr[left], arr[j] = arr[j], arr[left]
	return j
}

// quickSort函数用于递归地进行快速排序
func quickSort1(arr []int, left, right int) {
	// 子数组长度为 1 时终止递归
	if left >= right {
		return
	}

	p := partition(arr, left, right) // 获取分区点
	quickSort1(arr, left, p-1)       // 递归排序左边部分
	quickSort1(arr, p+1, right)      // 递归排序右边部分
}

func quickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	left, right := 0, len(arr)-1

	// 选择基准元素，这里选择中间的元素
	pivot := arr[len(arr)/2]

	// 分区过程
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// 递归对左右子数组进行快速排序
	if right > 0 {
		quickSort2(arr[:right+1])
	}
	if left < len(arr) {
		quickSort2(arr[left:])
	}
}

func test3() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("排序前:", arr)
	quickSort2(arr)
	fmt.Println("排序后:", arr)
}

func test4() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("排序前:", arr)
	quickSort1(arr, 0, len(arr)-1)
	fmt.Println("排序后:", arr)
}

func main() {
	// test3()
	test4()
}
