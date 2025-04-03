package main

/**
 * @Author nico
 * @Date 2025-04-03
 * @File: tool.go
 * @Description:
 */

import (
	"fmt"
	"reflect"
)

// 判断某一个值是否含在切片之中
func IsInArray[T comparable](item T, arr []T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

// ReverseArray 数组逆序
func ReverseArray[T any](arr []T) []T {
	// 创建一个新的切片来存放反转后的数据
	reversed := make([]T, len(arr))

	left, right := 0, len(arr)-1
	for left <= right {
		reversed[left], reversed[right] = arr[right], arr[left]
		left++
		right--
	}

	return reversed
}

// 数组的交集
func Intersect[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if IsInArray(v, b) {
			set = append(set, v)
		}
	}

	return set
}

// Chunk
// 按照指定的 groupSize 将一个切片分组，返回一个二维切片
func Chunk[T any](elements []T, groupSize int) [][]T {
	var result [][]T

	// 计算分组的数量
	groupCount := (len(elements) + groupSize - 1) / groupSize

	// 遍历原始切片，按 groupSize 大小进行分组
	for i := 0; i < groupCount; i++ {
		// 计算每组的结束索引，确保不超过原始切片的长度
		start := i * groupSize
		end := start + groupSize
		if end > len(elements) {
			end = len(elements)
		}

		// 将每组的切片添加到结果中
		result = append(result, elements[start:end])
	}

	return result
}

// 根据条件过滤对象数组
func Filter[T any](items []T, condition func(T) bool) []T {
	var result []T
	for _, item := range items {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

// 提取对象数组中某个字段的值，返回该字段值的数组
func Map[T any](objects []T, fieldName string) ([]interface{}, error) {
	var result []interface{}
	for _, obj := range objects {
		// 使用反射获取字段值
		v := reflect.ValueOf(obj)
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("field %s not found", fieldName)
		}
		result = append(result, field.Interface()) // 将字段值添加到结果切片中
	}
	return result, nil
}

// 查找某元素在数组中的下标
func FindIndex[T comparable](arr []T, target T) (index int, flag bool) {
	for i, item := range arr {
		if item == target {
			flag = true
			index = i
			return
		}
	}
	return
}

// 返回数组元素到索引的映射
func ItemMapIndex[T comparable](arr []T) map[T]int {
	result := make(map[T]int)
	for i, v := range arr {
		result[v] = i
	}
	return result
}

func ToAnySlice[T any](s []T) []any {
	result := make([]any, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
