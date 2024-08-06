package main

/**
 * @Author elastic·H
 * @Date 2024-08-03
 * @File: findMaxSubarray.go
 * @Description:
 * 有两个数组，求他们的公共最大长度子数组，返回其最大长度
 * 如 num1 和 num2 分别为 [1, 2, 3, 2, 1], [3, 2, 1, 4, 7] 则最大子数组为[3,2,1]，最大长度为 3
 * 解决办法：动态规划问题
 */

import (
	"fmt"
	"strings"
)

func findMaxLengthCommonSubarray(num1, num2 []int) int {
	m, n := len(num1), len(num2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLength := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if num1[i-1] == num2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			}
		}
	}

	return maxLength
}

// ///////////////////////////////////////////////////////////////////////////////////////
// 采用 两个数组均转成字符串 如1_2_3_2_1 和 3_2_1_4_7 后，取长度小的那个字符串（即数组长度小的）作为模式子串a，较大的为b，由于结果肯定不大于min(a 长度, b长度) 。因此用a的子串调用KMP看b中是否存在就行了！同时注意计算a的子串时候，从最大长度开始直至长度为1，类似于for(i=长度;i>=0;i--)

func buildKMPTable(pattern string) []int {
	table := make([]int, len(pattern))
	j := 0
	for i := 1; i < len(pattern); i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = table[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		table[i] = j
	}
	return table
}

func kmpSearch(text, pattern string) bool {
	if len(pattern) == 0 {
		return true
	}
	table := buildKMPTable(pattern)
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pattern[j] {
			j = table[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return true
		}
	}
	return false
}

func findMaxLengthCommonSubarray2(num1, num2 []int) int {
	// 将数组转换为字符串
	str1 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(num1)), "_"), "[]")
	str2 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(num2)), "_"), "[]")

	// 确保 str1 是较短的字符串
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}

	elements := strings.Split(str1, "_")
	for length := len(elements); length > 0; length-- {
		for start := 0; start <= len(elements)-length; start++ {
			subStr := strings.Join(elements[start:start+length], "_")
			if kmpSearch(str2, subStr) {
				return length
			}
		}
	}

	return 0
}

// ///////////////////////////////////////////////////////////////////////////////////////

// 返回其最大长度子数组，而不是最大长度
func findMaxLengthCommonSubarray3(num1, num2 []int) []int {
	m, n := len(num1), len(num2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLength := 0
	endIndex := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if num1[i-1] == num2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endIndex = i - 1
				}
			}
		}
	}

	// 构建最大公共子数组
	result := make([]int, maxLength)
	for i := 0; i < maxLength; i++ {
		result[i] = num1[endIndex-maxLength+1+i]
	}

	return result
}

func main() {
	num1 := []int{1, 2, 3, 2, 1}
	num2 := []int{3, 2, 1, 4, 7}

	result := findMaxLengthCommonSubarray(num1, num2)
	// result := findMaxLengthCommonSubarray2(num1, num2)
	// result := findMaxLengthCommonSubarray3(num1, num2)
	fmt.Println(" ---->  ", result)
}
