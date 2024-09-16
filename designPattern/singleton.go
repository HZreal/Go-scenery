package main

/**
 * @Author elastic·H
 * @Date 2024-09-16
 * @File: singleton.go
 * @Description: 单例模式
 */

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println(s1 == s2) // true
}
