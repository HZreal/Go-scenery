package main

/**
 * @Author elastic·H
 * @Date 2024-05-02
 * @File: mapStructure.go
 * @Description:
 */

import (
	"fmt"
)

func m() {
	m := make(map[string]int)
	m["a"] = 10
	v, ok := m["a"]
	if ok == true {
		fmt.Println(v)
	}
	v2 := m["ac"]
	fmt.Println(v2)

	// map[int]string{}
	// map[string]int{}
	// map[*int]string{}
	// map[[5]int]bool{}
	// map[struct{x int, y string}]int{}
}

func main() {
	m()
}
