package main

/**
 * @Author huang
 * @Date 2024-06-08
 * @File: mergeSort.go
 * @Description:
 */

import (
	"fmt"
)

// merge函数用于合并两个有序数组
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 将剩余的元素添加到结果数组中
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// mergeSort函数用于递归地进行归并排序
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("排序前:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("排序后:", sortedArr)
}
