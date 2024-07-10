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

/**
布农过滤器的基本原理
布农过滤器的核心原理是使用多个哈希函数将元素映射到位数组（bit array）中的多个位置。具体步骤如下：
	1.	初始化一个大小为  m  的位数组，并将所有位设置为 0。
	2.	使用  k  个不同的哈希函数将元素映射到位数组的  k  个位置，并将这些位置的位设置为 1。
	3.	检测一个元素是否存在于集合中时，同样使用这  k  个哈希函数将该元素映射到位数组的  k  个位置。如果这  k  个位置的位全为 1，则认为该元素可能在集合中；否则认为该元素不在集合中。

布农过滤器的优缺点
优点：
	•	空间效率高，能够用较少的空间存储大集合的信息。
	•	插入和查询操作时间复杂度为 O(k)，其中 k 是哈希函数的数量。
缺点：
	•	存在一定的误报率，即可能会错误地认为某个不存在的元素在集合中。
	•	不支持删除操作（有变种支持，但复杂度较高）。

布农过滤器的应用场景
	•	缓存系统：用于快速判断某个数据是否在缓存中，如果布农过滤器认为数据不在缓存中，可以直接查询数据库。
	•	垃圾邮件过滤：用于快速判断邮件是否为垃圾邮件。
	•	网络爬虫：用于判断 URL 是否已被访问过，避免重复抓取。
	•	数据库查询优化：用于快速判断数据是否存在于数据库的某个表中，以减少不必要的查询操作。
*/

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
