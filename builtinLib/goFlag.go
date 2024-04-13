package main

import (
	"flag"
	"fmt"
	"github.com/jessevdk/go-flags"
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

/**
go-flags

go-flags相比标准库flag支持更丰富的数据类型：
	所有的基本类型（包括有符号整数int/int8/int16/int32/int64，无符号整数uint/uint8/uint16/uint32/uint64，浮点数float32/float64，布尔类型bool和字符串string）和它们的切片；
	map 类型。只支持键为string，值为基础类型的 map；
	函数类型。
*/

type Option struct {
	IntFlag        int            `short:"i" long:"int" description:"int flag value"`
	IntSlice       []int          `long:"intslice" description:"int slice flag value"`
	BoolFlag       bool           `long:"bool" description:"bool flag value"`
	BoolSlice      []bool         `long:"boolslice" description:"bool slice flag value"`
	FloatFlag      float64        `long:"float" description:"float64 flag value"`
	FloatSlice     []float64      `long:"floatslice" description:"float64 slice flag value"`
	StringFlag     string         `short:"s" long:"string" description:"string flag value"`
	StringSlice    []string       `long:"strslice" description:"string slice flag value"`
	PtrStringSlice []*string      `long:"pstrslice" description:"slice of pointer of string flag value"`
	Call           func(string)   `long:"call" description:"callback"`
	IntMap         map[string]int `long:"intmap" description:"A map from string to int"`
}

func goFlags() {
	var opt Option
	opt.Call = func(value string) {
		fmt.Println("in callback: ", value)
	}

	_, err := flags.Parse(&opt)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Printf("int flag: %v\n", opt.IntFlag)
	fmt.Printf("int slice flag: %v\n", opt.IntSlice)
	fmt.Printf("bool flag: %v\n", opt.BoolFlag)
	fmt.Printf("bool slice flag: %v\n", opt.BoolSlice)
	fmt.Printf("float flag: %v\n", opt.FloatFlag)
	fmt.Printf("float slice flag: %v\n", opt.FloatSlice)
	fmt.Printf("string flag: %v\n", opt.StringFlag)
	fmt.Printf("string slice flag: %v\n", opt.StringSlice)
	fmt.Println("slice of pointer of string flag: ")
	for i := 0; i < len(opt.PtrStringSlice); i++ {
		fmt.Printf("\t%d: %v\n", i, *opt.PtrStringSlice[i])
	}
	fmt.Printf("int map: %v\n", opt.IntMap)
}

func main() {
	// useOsArgs()
	// flagArgs()

	goFlags()
	// go build -o goFlag goFlag.go
	// ./goFlag --int 42 --intslice 1 --intslice 2 --intslice 3 --bool --boolslice true --boolslice false --boolslice true --float 3.14 --floatslice 1.1 --floatslice 2.2 --floatslice 3.3 --string "hello" --strslice "one" --strslice "two" --strslice "three" --intmap key1:100 --intmap key2:200
}
