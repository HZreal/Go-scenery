package alg

/**
 * @Author elastic·H
 * @Date 2024-06-05
 * @File: alg.go
 * @Description:
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
