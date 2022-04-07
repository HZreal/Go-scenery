// 声明当前所在包
package main

//一个文件夹下的所有文件必须使用同一个包名

import (
	"fmt"
	"math"
	"strings"
)

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

	//var ccc bool          // 默认 false
	//fmt.Println(ccc)

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
	//resStr := multiRowStr + "hello!"    // fmt.Sprintf
	//fmt.Println(resStr)

	splitRes := strings.Split(multiRowStr, "\n\t")
	fmt.Println(splitRes)

	isContains := strings.Contains(multiRowStr, "e") // 是否包含
	fmt.Println(isContains)

	isPrefix := strings.HasPrefix(multiRowStr, "line") // 前缀判断
	fmt.Println(isPrefix)

	isSuffix := strings.HasSuffix(multiRowStr, "line3\n") // 后缀判断
	fmt.Println(isSuffix)

	//strings.Index()
	//strings.LastIndex()                                            // 子串出现的位置
	joinStr := strings.Join(splitRes, "\n\t") // join操作
	fmt.Println(joinStr)

}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte 打印中文等Unicode乱码
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
	//要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
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
	//逻辑位运算
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

// if 语句
func ifCondition() {
	// if条件判断
	var _num_1 = 10
	var _num_2 = 5

	if _num_1 > _num_2 {
		fmt.Println(_num_1)
	} else {
		fmt.Println(_num_2)
	}

	// switch
	var marks = 90
	var grade string = "D"
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(grade)

	// type-switch
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

	// select
	//类似于 switch 语句，但是select会随机执行一个可运行的case 如果没有case可运行，它将阻塞，直到有case可运行
}

// for 循环
func forCirculation() {
	// For循环有3中形式，只有其中的一种使用分号
	//for init; condition; post { }
	//for condition { }
	//for { }
	//init： 一般为赋值表达式，给控制变量赋初值；
	//condition： 关系表达式或逻辑表达式，循环控制条件；
	//post： 一般为赋值表达式，给控制变量增量或减量。
	//for语句执行过程如下：
	//①先对表达式 init 赋初值；
	//②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

	//	循环
	//s := "abc"
	//for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
	//	println(s[i])
	//}
	//var sum = 0
	//for he := 1; he <= 10; he++ {
	//	sum += he
	//}
	//fmt.Println("\nsum-----", sum)

	// 循环嵌套
	//var i, j int
	//for i = 2; i < 100; i++ {
	//	for j = 2; j <= (i / j); j++ {
	//		if i%j == 0 {
	//			break // 如果发现因子，则不是素数
	//		}
	//	}
	//	if j > (i / j) {
	//		fmt.Printf("%d  是素数\n", i)
	//	}
	//}

	// 无限循环
	//for true  {
	//	fmt.Printf("这是无限循环。\n");
	//}

	// range  类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
	//可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	//for key, value := range oldMap {
	//	newMap[key] = value
	//}
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
	//主要是使用场景不同
	//for可以
	//遍历array和slice
	//遍历key为整型递增的map
	//遍历string
	//for range可以完成所有for可以做的事情，却能做到for不能做的，包括
	//遍历key为string类型的map并同时获取key和value
	//遍历channel

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

	//defineVariable()
	//stringOperation()
	//traversalString()
	//changeString()
	//bitOperation()
	//ifCondition()
	//forCirculation()
	dataTypeConvert()

}
