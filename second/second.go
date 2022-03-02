package main

import "fmt"

func main() {

	// 数组
	var arr [10]float32
	var balance1 = [5]float32{1.0, 2.1, 3.0, 4.0, 5.0}
	balance2 := [3]float32{1.0, 2.1, 3.0}
	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(arr, balance1, balance2, balance3)

	// 指针
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &a        /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是: %x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 指针数组(存指针的数组，数组元素为指针)
	const MAX int = 3
	b := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &b[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
