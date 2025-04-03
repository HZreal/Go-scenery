package main

/**
 * @Author nico
 * @Date 2025-04-03
 * @File: tool.go
 * @Description:
 */

import (
	"fmt"
)

func t() {
	fmt.Println("t")
}

// 通过一个 map 将当前 map 映射到一个新 map
func MapByMap[T comparable](current map[T]string, by map[T]T) map[T]string {
	mm := make(map[T]string)
	for k, v := range current {
		mm[by[k]] = v
	}

	return mm
}
