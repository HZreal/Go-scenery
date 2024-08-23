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

/**
 * 无重复字符的最长子串
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters
 */
func unDuplicatedCharacterSubstring(str string) int {
	strLen := len(str)
	if strLen == 0 || strLen == 1 {
		return strLen
	}

	// strLen > 1
	res := 1
	for i := 0; i < strLen; i++ {
		tmpMap := map[byte]int{}
		tmpRes := 1
		tmpMap[str[i]]++

		for j := i + 1; j < strLen; j++ {
			// _, ok := tmpMap[str[j]]
			// if ok {
			// 	break
			// } else {
			// 	tmpRes++
			// 	tmpMap[str[j]] = 1
			// }

			tmpMap[str[j]]++
			if tmpMap[str[j]] > 1 {
				break
			}
			tmpRes++

		}
		if tmpRes > res {
			res = tmpRes
		}

	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

/**
 * 反转链表
 * https://leetcode.cn/problems/reverse-linked-list/description/
 */

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func reverseLinkedList(head *LinkedNode) *LinkedNode {
	// 定义一个节点，用于存储遍历到 curr 时的上一个节点
	var prev *LinkedNode
	curr := head
	for curr != nil {
		// 暂存 next 节点
		next := curr.Next
		// 修改当前节点的 next 指针
		curr.Next = prev
		// 更新 prev，用于下次迭代
		prev = curr
		// 更新当前节点
		curr = next
	}
	return prev
}

func testReverseLinkedList(arr []int) (res []int) {
	// 数组转单链表
	if len(arr) == 0 {
		return res
	}

	head := &LinkedNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &LinkedNode{Val: arr[i]}
		current = current.Next
	}

	// 反转单链表测试
	newHead := reverseLinkedList(head)

	// 将反转后的单链表转成数组
	current2 := newHead
	for current2 != nil {
		res = append(res, current2.Val)
		current2 = current2.Next
	}

	return res
}
