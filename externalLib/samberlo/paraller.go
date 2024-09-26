package main

import (
	"fmt"
	lo "github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"strconv"
	"strings"
)

/*
*
工具库 samber/lo
类似于 lodash.js
github地址：https://github.com/samber/lo
*/

func t1() {
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	fmt.Println(names)
	// []string{"Samuel", "John"}

	even := lo.Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0
	})
	fmt.Println(even)
	// []int{2, 4}

	aa := lo.Map([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Println(aa)
	// []string{"1", "2", "3", "4"}

	matching := lo.FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})
	fmt.Println(matching)
	// []string{"xpu", "xpu"}

	// bb := lo.FlatMap([]int{0, 1, 2}, func(x int, _ int) []string {
	// 	return []string{
	// 		strconv.FormatInt(x, 10),
	// 		strconv.FormatInt(x, 10),
	// 	}
	// })
	// fmt.Println(bb)
	// []string{"0", "0", "1", "1", "2", "2"}

	sum := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	fmt.Println(sum)
	// 10

}

func t2() {
	res := lop.Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Println(res)
	// []string{"1", "2", "3", "4"}

}

func taskGroupInBatch() {
	// 假设有 100 个任务
	tasks := make([]int, 100)
	for i := 0; i < 100; i++ {
		tasks[i] = i + 1 // 任务编号 1 到 100
	}

	// 使用 lo.Chunk 将任务分成每组 20 个
	chunkedTasks := lo.Chunk(tasks, 20)

	// 打印每一组任务
	for i, group := range chunkedTasks {
		fmt.Printf("第 %d 组任务: %v\n", i+1, group)
	}
}

func main() {
	// t1()
	// t2()

	taskGroupInBatch()
}
