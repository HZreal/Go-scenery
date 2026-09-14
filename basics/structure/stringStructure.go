package main

/**
 * @Author elastic·H
 * @Date 2024-05-01
 * @File: stringStructure.go
 * @Description:
 */

import (
	"fmt"
)

func t() {
	fmt.Println("t")

	// strings.Builder{}

	a := make([]int, 3, 5)
	fmt.Println(a)
	a[0] = 2
	fmt.Println(a)
}

func main() {
	t()
}
