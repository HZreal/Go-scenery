package main

/**
 * @Author elastic·H
 * @Date 2024-09-24
 * @File: stringJson.go
 * @Description: 字符串形式的 JSON
 */

import (
	"encoding/json"
	"fmt"
)

func marshalStringToJsonString() {
	// Go 中的一个普通字符串
	str := "Hello, World!"

	// 序列化为 JSON 格式的字符串
	jsonData, err := json.Marshal(str)
	if err != nil {
		fmt.Println("序列化错误:", err)
		return
	}

	// 输出 JSON 字符串
	fmt.Println(string(jsonData)) // 输出："Hello, World!"
}

func unmarshalJsonStringToString() {
	// JSON 格式的字符串（带引号）
	jsonData := `"Hello, World!"`

	// 定义一个 Go 的字符串变量
	var str string

	// 反序列化 JSON 字符串为 Go 字符串
	err := json.Unmarshal([]byte(jsonData), &str)
	if err != nil {
		fmt.Println("反序列化错误:", err)
		return
	}

	// 输出 Go 的字符串
	fmt.Println(str) // 输出：Hello, World!
}

func main() {
	marshalStringToJsonString()
	unmarshalJsonStringToString()
}
