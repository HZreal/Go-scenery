package main

/**
 * @Author huang
 * @Date 2024-07-11
 * @File: bloomFilter.go
 * @Description:
 */

import (
	"fmt"
	"hash/fnv"
)

// BloomFilter 布农过滤器结构
type BloomFilter struct {
	bitset []bool
	k      int
	m      int
}

// NewBloomFilter 初始化布农过滤器
func NewBloomFilter(m int, k int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, m),
		k:      k,
		m:      m,
	}
}

// hash 哈希函数
func hash(data string, seed int) int {
	h := fnv.New32a()
	h.Write([]byte(data))
	return int(h.Sum32()) + seed
}

// Add 向布农过滤器中添加元素
func (bf *BloomFilter) Add(item string) {
	for i := 0; i < bf.k; i++ {
		index := hash(item, i) % bf.m
		bf.bitset[index] = true
	}
}

// Contains 检查元素是否可能在布农过滤器中
func (bf *BloomFilter) Contains(item string) bool {
	for i := 0; i < bf.k; i++ {
		index := hash(item, i) % bf.m
		if !bf.bitset[index] {
			return false
		}
	}
	return true
}

func main() {
	m := 1000 // 位数组大小
	k := 3    // 哈希函数数量

	bf := NewBloomFilter(m, k)

	// 添加元素
	bf.Add("hello")
	bf.Add("world")

	// 检查元素
	fmt.Println(bf.Contains("hello"))  // 输出: true
	fmt.Println(bf.Contains("world"))  // 输出: true
	fmt.Println(bf.Contains("golang")) // 输出: false (可能会误报为 true)
}
