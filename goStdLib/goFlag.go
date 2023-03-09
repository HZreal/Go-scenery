package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

/*
os.Args获取命令行参数，是一个[]string，它的第一个元素是执行文件的名称。
os.Args是一个存储命令行参数的字符串切片，
*/
func useOsArgs() {
	argsSlice := os.Args
	if len(argsSlice) > 0 {
		for index, args := range argsSlice {
			fmt.Printf("args[%d]=%v\n", index, args)
		}
	}
}

/*
 */
func flagArgs() {
	// 两种常用的定义命令行flag参数的方法。
	// 1. flag.Type(flagName, value, desc) *Type
	// name1 := flag.String("name", "张三", "姓名")
	// age1 := flag.Int("age", 18, "年龄")
	// married1 := flag.Bool("married", false, "婚否")
	// delay1 := flag.Duration("d", 0, "时间间隔")

	// 2. flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	// 定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。
	flag.Parse()
	// 支持的命令行参数格式有以下几种：
	// 	-flag xxx （使用空格，一个-符号）
	// 	--flag xxx （使用空格，两个-符号）
	// 	-flag=xxx （使用等号，一个-符号）
	// 	--flag=xxx （使用等号，两个-符号）

	// Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。
	fmt.Println(name, age, married, delay)
	// 返回命令行参数后的其他参数，以[]string类型
	flag.Args()
	// 返回命令行参数后的其他参数个数
	flag.NArg()
	// 返回使用的命令行参数个数
	flag.NFlag()
}

func main() {
	// useOsArgs()
	flagArgs()
}
