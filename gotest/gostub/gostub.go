package main

/**
 * @Author elastic·H
 * @Date 2024-08-07
 * @File: gostub.go
 * @Description: gostub 测试框架
 * 四大测试框架: gostub goconvey gomock monkey
 * gostub 主要用来给变量、函数、过程打桩 但是给函数打桩时，需要做侵入式修改
 * convey 主要用途是用来组织测试用例的，提供了很多断言，兼容go test，有web ui，保存代码可自动跑测试
 * gomock 主要用来给接口打桩的。mockgen可以生成对应的接口测试文件
 * monkey 主要也是用来给变量、函数打桩的
 */

import (
	"fmt"
	"github.com/prashantv/gostub"
	"testing"
)

// FetchDataFromServer 实际函数
func FetchDataFromServer() string {
	// 假设这个函数从服务器获取数据
	return "real data"
}

func GetData() string {
	return FetchDataFromServer()
}

// TestGetData 测试函数
func TestGetData(t *testing.T) {
	// 使用 gostub 创建存根
	stubs := gostub.New()
	defer stubs.Reset() // 确保测试结束时重置

	// 替换 FetchDataFromServer 的实现
	stubs.StubFunc(FetchDataFromServer, "stubbed data")

	// 调用被测试函数
	data := GetData()
	if data != "stubbed data" {
		t.Errorf("expected 'stubbed data', got '%s'", data)
	}
}

func main() {
	fmt.Println(GetData()) // 输出 "real data"
}
