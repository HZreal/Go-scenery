package main

import (
	"fmt"
	"goBasics/fir_package"
	"math"
)

// 类型相同的不定参数
func funcParamsWithUncertianLength(s string, args ...int) {
	fmt.Println("参数s-------> ", s)
	fmt.Println("其他不定长参数args-------> ", args)
	var sum int
	for _, v := range args {
		sum += v
	}
	fmt.Printf(s, sum)
}

// 类型不同的不定参数
func funcParamsWithUncertianLength2(args ...interface{}) { // 空接口类型，表示参数args有不同的参数类型
	fmt.Println("不定长参数args-------> ", args)
	for _, v := range args {
		fmt.Printf("type is %T, value is %v\n", v, v)
	}
}

// 函数空返回
func emptyReturnFunc() {
	fmt.Println("this is a function with empty return")
	return
}

// 命名返回参数(函数返回值定义了接收的变量名)
func returnFunc1(a, b int) (c, d int) { // 返回值指定了变量c
	c = a + b
	d = a - b
	return // return后未指定返回什么，则返回定义中指定返回的c
	//return 22, 12                              // return后有返回值，就使用return后面的返回值
}

// 命名返回参数(隐式返回、显式返回)
func returnFunc2(a, b int) (c int) { // 返回值指定了变量c
	//c = a + b
	//return                                  	 // 隐式返回

	// 当重新定义局部变量c时，必须显式返回
	//var c = a + b      // 重定义c不能在当前级别的位置
	{
		var c = a + b
		//return            // 报错
		return c // 必须显示(指定c)返回
	}
}

// 使用 defer 延迟调用
func returnFunc3(a, b int) (c int) {
	defer func() {
		c += 10
	}()
	c = a + b
	return // 执行c=a+b完后隐式返回：先执行defer即c+=10, 然后返回c
}
func returnFunc4(a, b int) (c int) {
	defer func() {
		fmt.Println("-------此时的c值为c=a+b+5---------", c)
		c += 10
	}()
	c = a + b
	return c + 5 // 执行c=a+b完后显示返回：先执行c=c+5, 然后执行defer, 最后返回
}

// 函数
func funBasics() {
	//1.
	//函数声明包含一个函数名，参数列表， 返回值列表和函数体，函数可以没有参数或接受多个参数
	//函数是第一类对象，可作为参数传递。建议将复杂签名定义为函数类型
	//没有函数体的函数声明，则该函数不是以Go实现的。通常以汇编语言实现

	// 2. 参数
	//值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
	//引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
	//默认情况，Go语言使用的是值传递，即在调用过程中不会影响到实际参数
	//注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
	//注意2：map、slice、chan、指针、interface默认以引用的方式传递

	//同类型的不定传参
	//Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个
	//在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可
	//func add(args ...int) {    //0个或多个参数
	//}
	//func add(a int, args…int) int {    //1个或多个参数
	//}
	//func add(a int, b int, args…int) int {    //2个或多个参数
	//}
	//注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数

	//不同类型的不定传参： 就是函数的参数和每个参数的类型都不是固定的。
	//用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的
	//func add(args ...interface{}) {
	//}

	//3. 返回值
	// "_"标识符，用来忽略函数的某个返回值
	// 空返回  如函数emptyReturnFunc， returnFunc1
	// 返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略
	// 			s := make([]int, 2)
	// 			s = test()   		// 报错: multiple-value test() in single-value context
	//			x, _ := test()   	// 正确
	// 多返回值可直接作为其他函数调用实参
	//funcParamsWithUncertianLength2(returnFunc1(20, 10))

	// 命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回  如函数 returnFunc1、returnFunc2
	// 命名返回参数可被同名局部变量遮蔽，此时必须显式返回 如函数 returnFunc2
	//命名返回参数允许 defer 延迟调用通过闭包读取和修改 如函数 returnFunc3
	//ret3 := returnFunc3(20, 13)
	//fmt.Println(ret3)
	//显式 return 返回前，会先修改命名返回参数 如函数 returnFunc4
	//ret4 := returnFunc4(20, 13)
	//fmt.Println(ret4)

	// 只要声明函数的返回值变量名称，就会在函数初始化时候为之赋值为0，而且在函数体作用域可见
	// 如 returnFunc2定义了返回变量c int，则函数初始化时c被赋予int型的零值，函数域内可用

}

// 自定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string { // 函数作为参数
	return fn(s, x, y)
}

// 匿名函数
func anonymousFunc() {
	//函数可以像普通变量一样被传递或使用
	//匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
	//匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))

	// 函数集(数组，元素为函数)
	funcArr := []func(x int) int{
		func(x int) int {
			return x + 10
		},
		func(x int) int {
			return x * 2
		},
		func(x int) int {
			return x << 2
		},
	}
	println(funcArr[2](2))

	// 函数作为结构体字段
	t := struct {
		name      string
		funcField func(a int) int
	}{
		name: "hh",
		funcField: func(a int) int {
			return a << 2
		},
	}
	println(t.funcField(2))

	// channel里传送
	// TODO --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())

}

// 函数调用
func funcCalled() {
	// 函数参数
	// 在默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
	// 注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
	// 注意2：map、slice、chan、指针、interface默认以引用的方式传递

	resNum := numMax(1, 6)
	fmt.Println("resNum---------------", resNum)
	str1, str2 := swap5("happy", "nice a day")
	fmt.Println("swap result-------------", str1, str2)

	//var addArr [] int
	addArr := []int{1, 2, 3}
	resArr := fir_package.AddValue(addArr, 5)
	fmt.Println(resArr)
}

// 闭包实现
func myClosure(a int) func() int { // 在此闭包中，可以认为 i := a + 10 是 inner 函数的环境
	i := a + 10
	inner := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return inner
}
func testClosure() func() { // 在汇编层 ，testClosure 实际返回的是 FuncVal{ func_address, closure_var_pointer ... }对象，其中包含了匿名函数地址、闭包对象指针。当调 匿名函数时，只需以某个寄存器传递该对象即可
	a := 10
	fmt.Printf("a (%p) = %d\n", &a, a)
	return func() {
		fmt.Printf("a (%p) = %d\n", &a, a)
	}
}

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// 返回2个闭包
func returnTwoClosure(base int) (func(i int) int, func(j int) int) {
	base += 100
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(j int) int {
		base -= j
		return base
	}
	return add, sub
}

// 闭包
func closureBasics() {
	// 闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)
	// 闭包可以用来完成信息隐藏
	//c := myClosure(5) 		     // 创建闭包，传入a=5，得到i=15这个环境，即为一个闭包环境c
	//c()            			   	 // 闭包调用，闭包内的i=16
	//c()            				 // 闭包调用，闭包内的i=17

	// 闭包复制的是原对象指针，这就很容易解释延迟引用现象
	//tC := testClosure()
	//tC()

	// 外部引用函数参数局部变量
	//tmp1 := add(10)
	//fmt.Println(tmp1(1), tmp1(2))
	//tmp2 := add(100)    		   // 此时tmp1和tmp2不是一个实体了
	//fmt.Println(tmp2(1), tmp2(2))

	// 返回2个闭包
	rtc1, rtc2 := returnTwoClosure(50) // 创建闭包
	println(rtc1(10), rtc2(25))        // base值变化：  100 -> 150 -> 160 -> 135
	// 此时base=135
	println(rtc1(15), rtc2(12)) // base值变化：  135 -> 150 -> 138

}

// 数字阶乘  5! = 5*4*3*2*1
func factorial(i int) int {
	if i == 1 {
		return 1
	}
	return i * factorial(i-1)
}

// 斐波那契数列(Fibonacci)   0 1 1 2 3 5 8 13
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

// 递归
func recursiveFuncBasic() {
	// 求阶乘
	//println(factorial(5))

	// 输出fibonaci数列
	for i := 0; i < 10; i++ {
		println(fibonaci(i))
	}
}

// 延迟调用defer
func deferBasics() {
	//defer特性：
	//1. 关键字 defer 用于注册延迟调用。
	//2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
	//3. 多个defer语句，按先进后出的方式执行。
	//4. defer语句中的变量，在defer声明时就决定了。
	//defer用途：
	//1. 关闭文件句柄
	//2. 锁资源释放
	//3. 数据库连接释放

	//1. 每个goroutine都维护一个自己的defer链表。
	//2. 新注册的defer会被添加到链表头。
	//3. defer链表执行时，从链表头开始执行。所以表现出倒叙执行。
	//4. 函数如果注册了defer函数，编译器会在代码底部插入deferreturn函数。
	//5. 函数执行到deferreturn时，会根据defer结构体中的字段判断当前链表头的defer是不是自己注册的，是则执行并删除，反之，代表当前函数注册的defer已经执行完了，函数结束。
	//6. go1.12之前(含)使用上述方法执行defer，有一下几个问题:
	//			_defer结构体在堆上分配
	//			需要操作defer链表
	//			defer参数需要在堆和栈之间相互拷贝，这导致了defer函数执行效率低，速度慢
	//7. 1.14版本直接将defer的代码展开，插入到父函数中，避免了在堆上分配_defer结构体，也不用操作链表，性能大幅提升。
	//8. 1.14之后，发生panic时无法通过defer链表找到展开的defer，所以_defer结构体又增加了几个字段，借助这些信息，panic处理流程可以通过栈扫描的方式找到这些没有被注册到defer链表的defer函数，并按照正确的顺序执行

	// return之后的语句先执行，defer后的语句后执行
	// 即return后面的语句执行完，再去执行defer，最后才真正的函数返回，所以defer依然可以修改本应该返回的结果!
	//详看  returnFunc3  returnFunc4

	//defer 执行顺序是先进后出   这个很自然,后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了
	//var whatever [5]struct{}
	//for i, _ := range whatever {
	//	defer fmt.Println(i)                  // 由于先进后出，输出顺序为 4,3,2,1,0
	//}

	// defer 碰上闭包
	var whatever2 [5]struct{}
	for i, _ := range whatever2 {
		f := func() { fmt.Println(i) } // 闭包
		defer f()                      // 函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4
	}

	//

}

func main() {

	//funBasics()
	//funcParamsWithUncertianLength("sum: %d", 1, 2, 3, 4, 5)
	//_slice := []int{1, 2, 3, 4, 5}
	//funcParamsWithUncertianLength("sum: %d", _slice...)    // 注意：使用 slice 对象做变参时，必须展开
	//funcParamsWithUncertianLength2("first params", 12, "third params", 32, [2]int{1, 2}, map[string]string{"name": "hu", "age": "22"})
	//anonymousFunc()
	//funcCalled()
	//dataTypeConvert()

	//closureBasics()
	//recursiveFuncBasic()

	deferBasics()

}

func numMax(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap5(x, y string) (string, string) { // 字符串互换
	return y, x
}

func swap2(x, y *int) { // 两个数据值交换
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
