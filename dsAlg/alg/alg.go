package alg

/**
 * @Author elastic·H
 * @Date 2024-06-05
 * @File: alg.go
 * @Description:
 */

/*
 * 二分查找
 * 双指针
 */
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] < target { // [left, mid, target, right]
			left = mid + 1
		} else if arr[mid] > target { // [left, target, mid, right]
			//
			right = mid - 1
		} else {
			//
			return mid
		}

	}
	return -1
}

/**
 * 移除元素
 * leetcode: https://leetcode.cn/problems/remove-element
 * 双指针，不能使用新数组则只能修改原数组
 */
func removeElement(nums []int, val int) int {
	// left := 0
	// for _, v := range nums {
	// 	if v != val {
	// 		nums[left] = v
	// 		left++
	// 	}
	// }
	// return left

	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slow] = val
			slow++
		}
	}
	return slow
}

// 双指针优化：依然使用双指针，两个指针初始时分别位于数组的首尾，向中间移动遍历该序列。
// 如果左指针 left 指向的元素等于 val，此时将右指针 right 指向的元素复制到左指针 left 的位置，然后右指针 right 左移一位。如果赋值过来的元素恰好也等于 val，可以继续把右指针 right 指向的元素的值赋值过来（左指针 left 指向的等于 val 的元素的位置继续被覆盖），直到左指针指向的元素的值不等于 val 为止。
// 与方法一不同的是，方法二避免了需要保留的元素的重复赋值操作。
func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

/**
 * 有序数组的平方
 * leetcode: https://leetcode.cn/problems/squares-of-a-sorted-array/
 * 双指针
 */
func sortedSquares(arr []int) []int {
	left, right := 0, len(arr)-1
	var res []int
	for left <= right {
		leftLeft := arr[left] * arr[left]
		rightRight := arr[right] * arr[right]
		if leftLeft <= rightRight {
			//
			res = append(res, rightRight)
			right--
		} else {
			//
			res = append(res, leftLeft)
			left++
		}
	}

	return reverse(res) // 自己实现逆转
	// return lo.Reverse(res)
}

// 数组逆序
func reverse(arr []int) []int {
	//
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func sortedSquares2(arr []int) []int {
	left, right := 0, len(arr)-1
	res2 := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {

		leftLeft := arr[left] * arr[left]
		rightRight := arr[right] * arr[right]

		if leftLeft <= rightRight {
			res2[i] = rightRight
			right--
		} else {
			res2[i] = leftLeft
			left++
		}
	}

	return res2
}
