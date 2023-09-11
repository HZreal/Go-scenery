// 声明当前所在包
package main

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"
)

// 一个文件夹下的所有文件必须使用同一个包名

// 导包顺序
// 	先从项目根目录的 vendor 目录中查找
// 	最后从 $GOROOT/src 目录下查找
// 	然后从 $GOPATH/src 目录下查找
// 	都找不到的话，就报错。

// 基本数据类型
func dataType() {
	//  int8  uint8类型大小为 1 字节，  表示十进制数-128 ~ 127, 0 ~ 255，         byte是uint8 的别名
	// int16 uint16类型大小为 2 字节，  表示十进制数-32768 ~ 32767, 0 ~ 65535，
	// int32 uint32类型大小为 4 字节，  表示十进制数-21亿 ~ 21亿, 0 ~ 42亿，       rune是int32 的别名
	// int64 uint64类型大小为 8 字节，  表示十进制数 ...
	// int类型的大小是和操作系统位数相关的，如果是32位操作系统，int 类型的大小就是4字节。如果是64位操作系统，int 类型的大小就是8个字节

	// 1.1.1. 整型
	// 整型分为以下两个大类：
	// 按长度分为：int8、int16、int32、int64    对应的无符号整型：uint8、uint16、uint32、uint64
	// 其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。

	// 1.1.2. 浮点型
	// Go语言支持两种浮点型数：float32和float64。这两种浮点型数据格式遵循IEEE 754标准：
	// float32 的浮点数的最大范围约为3.4e38，可以使用常量定义：math.MaxFloat32。
	// float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64。

}

// 定义变量
func defineVariable() {
	var username = "huang"
	fmt.Println(username)

	var password string
	password = "zhen"
	fmt.Println(password)

	var character1 = 'a' // 声明字符
	var character2 = '中' // 声明字符
	fmt.Println(character1, character2)

	var res string
	res = username + password
	fmt.Println(res)

	// var ccc bool          // 默认 false
	// fmt.Println(ccc)

	const WIDTH = 20
	const LENGTH int = 10
	const (
		PI = 3.14
		_n = iota
	)
	// iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

}

// 字符串常用操作
func stringOperation() {
	var multiRowStr string = `line1
	line2
	line3
	`
	fmt.Println(multiRowStr)
	// resStr := multiRowStr + "hello!"    // fmt.Sprintf
	// fmt.Println(resStr)

	splitRes := strings.Split(multiRowStr, "\n\t")
	fmt.Println(splitRes)

	isContains := strings.Contains(multiRowStr, "e") // 是否包含
	fmt.Println(isContains)

	isPrefix := strings.HasPrefix(multiRowStr, "line") // 前缀判断
	fmt.Println(isPrefix)

	isSuffix := strings.HasSuffix(multiRowStr, "line3\n") // 后缀判断
	fmt.Println(isSuffix)

	// strings.Index()
	// strings.LastIndex()                                            // 子串出现的位置
	joinStr := strings.Join(splitRes, "\n\t") // join操作
	fmt.Println(joinStr)

}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { // byte 打印中文等Unicode乱码
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { // rune类型(实际是一个int32)，  Go使用了特殊的 rune 类型来处理 Unicode
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 修改字符串
func changeString() {
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	fmt.Println(byteS1) // [104 101 108 108 111]
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "你好"
	runeS2 := []rune(s2)
	runeS2[0] = '他'
	fmt.Println(string(runeS2))
}

// 位运算
func bitOperation() {
	// 逻辑位运算
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)

}

func assertTypeSwitch() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
func useSelect() {
	var chan1, chan2, chan3 chan int
	var i1, i2 int

	select {
	case i1 = <-chan1:
		fmt.Printf("received ", i1, " from c1\n")
	case chan2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := <-chan3:
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
func timeOutJudge() {
	var resChan = make(chan *http.Response)
	// 异步请求，并继续执行下面select
	go func() {
		res, _ := http.Get("https://www.baidu.com")
		resChan <- res
	}()
	select {
	case data := <-resChan:
		fmt.Println("data-----\n", *data)
	case af := <-time.After(time.Second * 3):
		fmt.Println("request time out-----", af)
	}
}
func exitsProgram() {
	// 主线程（协程）中如下：
	var shouldQuit = make(chan interface{})
	go func() {
		defer func() {
			if err := recover(); err != nil {
				shouldQuit <- err
			}
		}()
		// 此协程中，模拟运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
		panic("panic error!")
	}()
	select {
	case err := <-shouldQuit:
		// cleanUp()
		fmt.Println("通道中有err，输出---", err)
		return
	case af := <-time.After(time.Second * 3):
		fmt.Println("request time out-----", af)
	}
}
func checkChannelIsBlocked() {
	ch := make(chan int, 5)
	//...
	data := 0
	select {
	case ch <- data: // 尝试向管道发数据，成功则执行case，表示管道未满
		fmt.Println("send successfully, channel is not full", ch)
	default: // 执行了default则表示上述尝试发送失败，即管道已满
		fmt.Println("channel is full")
	}
	fmt.Println("final ---")

}

// if 语句
func ifCondition() {
	// if条件判断
	// var _num_1 = 10
	// var _num_2 = 5
	// if _num_1 > _num_2 {
	// 	fmt.Println(_num_1)
	// } else {
	// 	fmt.Println(_num_2)
	// }

	// switch 基本使用
	// var marks = 90
	// var grade string = "D"
	// switch marks {
	// case 90:
	// 	grade = "A"
	// case 80:
	// 	grade = "B"
	// case 50, 60, 70:
	// 	grade = "C"
	// default:
	// 	grade = "D"
	// }
	// fmt.Println(grade)

	// type-switch  interface{}类型断言
	// assertTypeSwitch()

	// select
	// 1.1.1. select 语句
	// 类似于 switch 语句，其用于通信的switch语句，每个case必须是一个通信操作，要么是发送要么是接收
	// Go会按顺序从头到尾评估每一个发送和接收的语句，然后执行一个可运行的case 如果没有case可运行，它将阻塞，直到有case可运行
	// 		每个case都必须是一个通信操作
	// 		所有channel表达式都会被求值
	// 		所有被发送的表达式都会被求值
	// 		如果任意某个通信可以进行，它就执行；其他被忽略。
	// 		如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
	// 		否则：
	// 		如果有default子句，则执行该语句。同时程序的执行会从select语句后的语句中恢复
	// 		如果没有default字句，select将阻塞，直到某个case可以运行；Go不会重新对channel或值进行求值。
	// useSelect()

	// select可以监听channel的数据流动
	// switch语句可以选择任何使用相等比较的条件，select由比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作
	// 		select { 		            	// 不停的在这里检测
	// 			case <-chan1 : 		        // 检测有没有数据可以读
	// 				// 如果chan1成功读取到数据，则进行该case处理语句
	// 			case chan2 <- 1 : 		    // 检测有没有可以写
	// 				// 如果成功向chan2写入数据，则进行该case处理语句
	//
	// 			...
	//
	// 			// 假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞//一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
	// 			default:
	// 				// 如果以上都没有符合条件，那么则进行default处理流程
	// 		}

	// 1.1.2. 典型用法
	// 1.超时判断
	// 比如在下面的场景中，使用全局resChan来接受response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行
	// timeOutJudge()

	// 2.退出
	// exitsProgram()

	// 3.判断channel是否阻塞
	// 某些情况下，不希望channel缓存满了
	checkChannelIsBlocked()

}

// for 循环
func forCirculation() {
	// For循环有3中形式，只有其中的一种使用分号
	// for init; condition; post { }
	// for condition { }
	// for { }
	// init： 一般为赋值表达式，给控制变量赋初值；
	// condition： 关系表达式或逻辑表达式，循环控制条件；
	// post： 一般为赋值表达式，给控制变量增量或减量。
	// for语句执行过程如下：
	// ①先对表达式 init 赋初值；
	// ②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

	//	循环
	// s := "abc"
	// for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
	//	println(s[i])
	// }
	// var sum = 0
	// for he := 1; he <= 10; he++ {
	//	sum += he
	// }
	// fmt.Println("\nsum-----", sum)

	// 循环嵌套
	// var i, j int
	// for i = 2; i < 100; i++ {
	//	for j = 2; j <= (i / j); j++ {
	//		if i%j == 0 {
	//			break // 如果发现因子，则不是素数
	//		}
	//	}
	//	if j > (i / j) {
	//		fmt.Printf("%d  是素数\n", i)
	//	}
	// }

	// 无限循环
	// for true  {
	//	fmt.Printf("这是无限循环。\n");
	// }

	// range  类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
	// 可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	// for key, value := range oldMap {
	//	newMap[key] = value
	// }
	// 注意:range 会复制对象，key, value都是从复制品中取出
	a := [3]int{0, 1, 2}
	for i, v := range a { // index、value 都是从复制品中取出。
		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}
		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
	}
	// 建议改用引用类型，其底层数据不会被复制
	ss := []int{1, 2, 3, 4, 5}
	for i, v := range ss { // 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			ss = ss[:3] // 对 slice 的修改，不会影响 range。
			ss[2] = 100 // 对底层数据的修改。
		}
		println(i, v)
	}
	// 另外两种引用类型 map、channel 是指针包装，而不像 slice 是 struct

	// for 和 for range有什么区别?
	// 主要是使用场景不同
	// for可以
	// 遍历array和slice
	// 遍历key为整型递增的map
	// 遍历string
	// for range可以完成所有for可以做的事情，却能做到for不能做的，包括
	// 遍历key为string类型的map并同时获取key和value
	// 遍历channel

	// Goto、Break、Continue
	// 1.三个语句都可以配合标签(label)使用
	// 2.标签名区分大小写，定以后若不使用会造成编译错误
	// 3.continue、break配合标签(label)可用于多层循环跳出
	// 4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同

	// goto语句
	var _aa int = 10
LOOP:
	for _aa < 20 {
		if _aa == 15 {
			/* 跳过迭代 */
			_aa = _aa + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", _aa)
		_aa++
	}

}

// 强制类型转换        Go语言只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用
func dataTypeConvert() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {

	// defineVariable()
	// stringOperation()
	// traversalString()
	// changeString()
	// bitOperation()
	ifCondition()
	// forCirculation()
	// dataTypeConvert()

}
