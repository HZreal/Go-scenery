package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

// Snowflake 包提供了以下功能：
// 一个非常简单的 Twitter 雪花生成器
// 解析现有雪花 ID 的方法
// 将雪花 ID 转换为多种其他数据类型并返回的方法
// JSON Marshal/Unmarshal 函数可在 JSON API 中轻松使用雪花 ID
// 单调时钟计算可防止时钟偏移
//
// 默认情况下，生成的 id 有以下格式：
// 整个 ID 是一个 63 位整数，存储在 int64 中
// 41 位用于存储毫秒精度的时间戳，使用自定义纪元
// 10 位用于存储节点 ID 范围从 0 - 1023
// 12 位用于存储序列号 范围从0 - 4095
// 因此，其生成的 int64 的 id 长度为 19 位。
// +--------------------------------------------------------------------------+
// | 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
// +--------------------------------------------------------------------------+

func testSnowflake1() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	fmt.Println("id ---->  ", int64(id))

	for i := 0; i < 10; i++ {
		// 生成连续的 id
		fmt.Println(int64(node.Generate()))
	}
}
func main() {
	testSnowflake1()
}

// Snowflake 库通常用于需要生成唯一 ID 的场景，例如分布式系统中的分布式 ID 生成器、消息队列中的消息 ID 生成器、分布式锁中的锁 ID 生成器等。在这些场景中，需要生成的 ID 必须是全局唯一的，以避免出现 ID 冲突的情况。
//
// Snowflake 库生成的 ID 是一个 64 位的整数，其中包含了时间戳、节点 ID 和序列号等信息，这些信息可以用来保证 ID 的唯一性。同时，Snowflake 库的生成速度非常快，可以在高并发的环境下快速生成大量的 ID。
