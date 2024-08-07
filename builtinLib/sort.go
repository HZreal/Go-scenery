package main

/**
 * @Author elastic·H
 * @Date 2024-08-07
 * @File: sort.go
 * @Description:
 */

import (
	"fmt"
	"sort"
)

func useCase1() {
	ints := []int{3, 1, 4, 1, 5, 9}
	sort.Ints(ints)
	fmt.Println(ints) // 输出: [1 1 3 4 5 9]
}

func useCase2() {
	floats := []float64{3.1, 2.2, 1.5}
	sort.Float64s(floats)
	fmt.Println(floats) // 输出: [1.5 2.2 3.1]
}

func useCase3() {
	strings := []string{"banana", "apple", "pear"}
	sort.Strings(strings)
	fmt.Println(strings) // 输出: [apple banana pear]
}

type Person struct {
	Name string
	Age  int
}

// 对象数组的排序
func useCase4() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people) // 输出: [{Bob 25} {Alice 30} {Charlie 35}]
}

type personList []Person

func (a personList) Len() int {
	return len(a)
}
func (a personList) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a personList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func useCase5() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Sort(personList(people))
	fmt.Println(people) // 输出: [{Bob 25} {Alice 30} {Charlie 35}]
}

func useCase6() {
	ints := []int{1, 3, 5, 7, 9}
	// 返回符合条件的最小的索引
	index := sort.Search(len(ints), func(i int) bool {
		return ints[i] >= 5
	})
	fmt.Println(index) // 输出: 2 （即 5 所在的索引）
}

func main() {
	// useCase1()
	// useCase2()
	// useCase3()
	// useCase4()
	// useCase5()
	useCase6()

}
