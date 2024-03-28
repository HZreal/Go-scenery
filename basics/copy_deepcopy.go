package main

import (
	"fmt"
	"reflect"
)

//浅拷贝：浅拷贝只复制数据的顶层结构，而不会递归复制数据中的引用类型数据。因此，如果原始数据包含引用类型的字段（如切片、映射、指针等），则浅拷贝将只复制它们的引用，而不会复制引用指向的实际数据。这意味着修改拷贝后的对象中的引用类型数据会影响原始对象中的数据。
//深拷贝：深拷贝会递归地复制所有的数据，包括引用类型的数据。这意味着创建的拷贝是完全独立于原始对象的，修改拷贝后的对象不会影响原始对象。

type Person struct {
	Name   string
	Age    int
	Colors []string
}

func testCopy() {
	// 原始对象
	p1 := Person{
		Name:   "Alice",
		Age:    30,
		Colors: []string{"red", "blue", "green"},
	}

	// 浅拷贝
	p2 := p1

	// 修改 p2 中的引用类型数据
	p2.Colors[0] = "yellow"

	// p1 中的引用类型数据也被修改了
	fmt.Println(p1.Colors) // 输出: [yellow blue green]
}

func deepCopy(src, dst interface{}) {
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcValue.Kind() != reflect.Ptr || dstValue.Kind() != reflect.Ptr {
		fmt.Println("参数必须是指针类型")
		return
	}

	if srcValue.Elem().Type() != dstValue.Elem().Type() {
		fmt.Println("源对象和目标对象类型不匹配")
		return
	}

	dstValue.Elem().Set(srcValue.Elem())
}

func testDeepCopy() {
	// 原始对象
	p1 := &Person{
		Name:   "Alice",
		Age:    30,
		Colors: []string{"red", "blue", "green"},
	}

	// 深拷贝
	var p2 Person
	deepCopy(p1, &p2)

	// 修改 p2 中的引用类型数据
	p2.Colors[0] = "yellow"

	// p1 中的引用类型数据不受影响
	fmt.Println(p1.Colors) // 输出: [red blue green]
}

func main() {
	// testCopy()

	// TODO 有问题
	testDeepCopy()
}
