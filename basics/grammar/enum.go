package main

/**
 * @Author nico
 * @Date 2024-12-23
 * @File: enum.go
 * @Description:
 */

import (
	"fmt"
)

func t() {
	fmt.Println("t")
}

//	---------------------------------------------------------------------------------
// 枚举一般从 1 开始

type Operation int

const (
	// Add Operation = iota  // bad
	Add Operation = iota + 1 // good, 枚举一般从 1 开始
	Subtract
	Multiply
)

// Add=1, Subtract=2, Multiply=3

// ---------------------------------------------------------------------------------
// 某些情况下，使用零值是有意义的（枚举从零开始），例如，当零值是理想的默认行为时

type LogOutput int

const (
	LogToStdout LogOutput = iota
	LogToFile
	LogToRemote
)

// LogToStdout=0, LogToFile=1, LogToRemote=2
