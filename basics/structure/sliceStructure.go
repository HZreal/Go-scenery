package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
 * @Author elastic·H
 * @Date 2024-05-01
 * @File: sliceStructure.go
 * @Description:
 */

func test(s []int) {
	printSliceStruct(&s)
}

func printSliceStruct(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("slice struct: %+v, slice is %v\n ", ss, s)
}

func sliceTest1() {
	s := make([]int, 3, 5)
	fmt.Printf("TTTT%T  vvvv%v", s, s)
	printSliceStruct(&s)
	test(s)
}

// /////////////////////////////////// 扩容机制 /////////////////////////////////////////////////
// 参考源码：runtime/slice.go
// 预估新长度 newLen 大于旧切片容量 oldCap 的 2 倍时，直接扩容到其新长度
// 反之（2 倍及以内），比较旧切片的容量 oldCap 与阈值 threshold（1.18 后为 256，之前为 1024） 的大小  ----- 保证平滑的过度
// 		若 oldCap < threshold，即对于小的切片，直接扩容到旧容量 oldCap 的 2 倍
// 		反之（oldCap >= threshold），即对于大的切片，采用旧容量 1.25 倍不断尝试递增，直至大小超过新长度 newLen
// 基于新容量的计算完成，再进行实际的内存分配

// ////////////////////////////////////////////////////////////////////////////////////

func doAppend(s []int) {
	s = append(s, 1)
	printLenAndCap(s)
}

func printLenAndCap(s []int) {
	fmt.Println(s)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
}
func sliceTest2() {
	// 字节一面
	s := make([]int, 8, 8)

	doAppend(s[:4])   // 传入子切片，append 时操作的是底层数组，未触发扩容
	printLenAndCap(s) // 可以看到原切片的信息(len、cap)不变，但底层数组被修改了
	doAppend(s)       // 传入子切片（值拷贝，只是看着还像是 s 而已），append 操作时发现需要扩容，新建底层数组，不影响原底层数组（即原底层数组的引用依然可用）
	printLenAndCap(s) // 可以看到原切片的信息(len、cap)不变，原因是底层数组未改变，上一步是新增底层数组
}

// ////////////////////////////////////////////////////////////////////////////////////
func sliceTest3() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)                  // 共用底层数组，修改了str2，str1 也会受影响
	str2 = append(str2, "z", "x", "y") // append 导致底层数组扩容，str2 引用新数组，而 str1 还是引用旧数组
	fmt.Println(str1, str2)
}

// ////////////////////////////////////////////////////////////////////////////////////

func main() {
	// sliceTest1()
	// sliceTest2()
	sliceTest3()
}
