package main

import (
	"errors"
	"fmt"
	"goBasics/fir_package" // 从 $GOPATH/src 或 $GOROOT/src 或者 $GOPATH/pkg/mod 目录下搜索包并导入
	// "../fir_package"                 // Go Modules 不支持相对导入
	"math"
	"net/http"
	"os"
	"sync"
	"time"
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
	// return 22, 12                              // return后有返回值，就使用return后面的返回值
}

// 命名返回参数(隐式返回、显式返回)
func returnFunc2(a, b int) (c int) { // 返回值指定了变量c
	// c = a + b
	// return                                  	 // 隐式返回

	// 当重新定义局部变量c时，必须显式返回
	// var c = a + b      // 重定义c不能在当前级别的位置
	{
		var c = a + b
		// return            // 报错
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
	// 1.
	// 函数声明包含一个函数名，参数列表， 返回值列表和函数体，函数可以没有参数或接受多个参数
	// 函数是第一类对象，可作为参数传递。建议将复杂签名定义为函数类型
	// 没有函数体的函数声明，则该函数不是以Go实现的。通常以汇编语言实现

	// 2. 参数
	// 值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
	// 引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
	// 默认情况，Go语言使用的是值传递，即在调用过程中不会影响到实际参数
	// 注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
	// 注意2：map、slice、chan、指针、interface默认以引用的方式传递

	// 同类型的不定传参
	// Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个
	// 在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可
	// func add(args ...int) {    //0个或多个参数
	// }
	// func add(a int, args…int) int {    //1个或多个参数
	// }
	// func add(a int, b int, args…int) int {    //2个或多个参数
	// }
	// 注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数

	// 不同类型的不定传参： 就是函数的参数和每个参数的类型都不是固定的。
	// 用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的
	// func add(args ...interface{}) {
	// }

	// 3. 返回值
	// "_"标识符，用来忽略函数的某个返回值
	// 空返回  如函数emptyReturnFunc， returnFunc1
	// 返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略
	// 			s := make([]int, 2)
	// 			s = test()   		// 报错: multiple-value test() in single-value context
	//			x, _ := test()   	// 正确
	// 多返回值可直接作为其他函数调用实参
	// funcParamsWithUncertianLength2(returnFunc1(20, 10))

	// 命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回  如函数 returnFunc1、returnFunc2
	// 命名返回参数可被同名局部变量遮蔽，此时必须显式返回 如函数 returnFunc2
	// 命名返回参数允许 defer 延迟调用通过闭包读取和修改 如函数 returnFunc3
	// ret3 := returnFunc3(20, 13)
	// fmt.Println(ret3)
	// 显式 return 返回前，会先修改命名返回参数 如函数 returnFunc4
	// ret4 := returnFunc4(20, 13)
	// fmt.Println(ret4)

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
	// 函数可以像普通变量一样被传递或使用
	// 匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
	// 匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送
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

	// var addArr [] int
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

// 一般闭包?
func commomClosure(i int) func(args ...interface{}) int {
	inner := func(args ...interface{}) int {
		fmt.Println(args)
		return i + 1
	}
	return inner
}

// 闭包
func closureBasics() {
	// 闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)
	// 闭包可以用来完成信息隐藏
	// c := myClosure(5) 		     // 创建闭包，传入a=5，得到i=15这个环境，即为一个闭包环境c
	// c()            			   	 // 闭包调用，闭包内的i=16
	// c()            				 // 闭包调用，闭包内的i=17

	// 闭包复制的是原对象指针，这就很容易解释延迟引用现象
	// tC := testClosure()
	// tC()

	// 外部引用函数参数局部变量
	// tmp1 := add(10)
	// fmt.Println(tmp1(1), tmp1(2))
	// tmp2 := add(100)    		   // 此时tmp1和tmp2不是一个实体了
	// fmt.Println(tmp2(1), tmp2(2))

	// 返回2个闭包
	rtc1, rtc2 := returnTwoClosure(50) // 创建闭包
	println(rtc1(10), rtc2(25))        // base值变化：  100 -> 150 -> 160 -> 135
	// 此时base=135
	println(rtc1(15), rtc2(12)) // base值变化：  135 -> 150 -> 138

	//
	// inner := commomClosure(2)
	// inner()

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
	// println(factorial(5))

	// 输出fibonaci数列
	for i := 0; i < 10; i++ {
		println(fibonaci(i))
	}
}

// 大量循环下 测试滥用defer
var lock sync.Mutex

func test1() {
	lock.Lock()
	lock.Unlock()
}
func test2() {
	lock.Lock()
	defer lock.Unlock()
}
func testWithoutDefer() {
	t1 := time.Now()

	for i := 0; i < 10000; i++ {
		test1()
	}
	elapsed := time.Since(t1)
	fmt.Println("test elapsed: ", elapsed)
}
func testWithDefer() {
	t1 := time.Now()

	for i := 0; i < 10000; i++ {
		test2()
	}
	elapsed := time.Since(t1)
	fmt.Println("testdefer elapsed: ", elapsed)
}

// defer 与 closure
func testDeferWithMulClosure(a, b int) (res int, err error) {

	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) {
		fmt.Printf("second defer err %v\n", err)
	}(err)
	defer func() {
		fmt.Printf("third defer err %v\n", err)
	}()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	res = a / b
	return
}

// defer 与 return
func testDeferWithReturn() (i int) {
	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2 // 在有具名返回值的函数中（这里具名返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。所以defer closure 输出结果为 2 而不是 1
}

// defer nil 函数
func testDeferWithNil() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

// 错误的位置使用 defer
func useDeferIncorrectly() error {
	res, err := http.Get("https://www.baidu.com")
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// ..code...

	return nil
}

// 正确使用 defer
func useDeferCorrectly() error {
	// 当有错误的时候，err 会被返回，否则当整个函数返回的时候，会关闭 res.Body

	// 在这里，你同样需要检查 res 的值是否为 nil ，这是 http.Get 中的一个警告。通常情况下，出错的时候，返回的内容应为空并且错误会被返回，可当你获得的是一个重定向 error 时， res 的值并不会为 nil ，但其又会将错误返回。上面的代码保证了无论如何 Body 都会被关闭，如果你没有打算使用其中的数据，那么你还需要丢弃已经接收的数据
	res, err := http.Get("https://www.baidu.com")
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	// ..code...

	return nil
}

// 延迟调用defer
func deferBasics() {
	// defer特性：
	// 1. 关键字 defer 用于注册延迟调用。
	// 2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
	// 3. 多个defer语句，按先进后出的方式执行。
	// 4. defer语句中的变量，在defer声明时就决定了。
	// defer用途：
	// 1. 关闭文件句柄
	// 2. 锁资源释放
	// 3. 数据库连接释放

	// 1. 每个goroutine都维护一个自己的defer链表。
	// 2. 新注册的defer会被添加到链表头。
	// 3. defer链表执行时，从链表头开始执行。所以表现出倒叙执行。
	// 4. 函数如果注册了defer函数，编译器会在代码底部插入deferreturn函数。
	// 5. 函数执行到deferreturn时，会根据defer结构体中的字段判断当前链表头的defer是不是自己注册的，是则执行并删除，反之，代表当前函数注册的defer已经执行完了，函数结束。
	// 6. go1.12之前(含)使用上述方法执行defer，有一下几个问题:
	//			_defer结构体在堆上分配
	//			需要操作defer链表
	//			defer参数需要在堆和栈之间相互拷贝，这导致了defer函数执行效率低，速度慢
	// 7. 1.14版本直接将defer的代码展开，插入到父函数中，避免了在堆上分配_defer结构体，也不用操作链表，性能大幅提升。
	// 8. 1.14之后，发生panic时无法通过defer链表找到展开的defer，所以_defer结构体又增加了几个字段，借助这些信息，panic处理流程可以通过栈扫描的方式找到这些没有被注册到defer链表的defer函数，并按照正确的顺序执行

	// return之后的语句先执行，defer后的语句后执行
	// 即return后面的语句执行完，再去执行defer，最后才真正的函数返回，所以defer依然可以修改本应该返回的结果!
	// 详看  returnFunc3  returnFunc4

	// defer 执行顺序是先进后出   这个很自然，后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了
	// var whatever [5]struct{}
	// for i, _ := range whatever {
	//	defer fmt.Println(i)                  // 每次循环时将defer操作放入队列，最后逆序执行，输出顺序为 4,3,2,1,0
	// }

	// defer 碰上闭包
	// var whatever2 [5]struct{}
	// for i, _ := range whatever2 {
	//	f := func() {                        // 闭包
	//		fmt.Println(i)                   // 由于闭包用到的变量 i 在逆序执行的时候已经变成4，所以输出全都是4
	//	}
	//	defer f()
	// }

	// 将循环变量i通过参数temp传入闭包(值拷贝)，则在defer逆序执行时，每个闭包中的变量参数因为值拷贝，都不一样了
	// var whatever3 [5]struct{}
	// for i, _ := range whatever3 {
	//	f := func(j int) {
	//		fmt.Println(i, j)               // 由于j是值拷贝传进闭包，i依然为循环变量 所以defer逆序执行的时候，闭包用到的参数j分别是4,3,2,1,0，i分别是4,4,4,4,4，所以输出第一列：4,4,4,4,4，第二列：4,3,2,1,0
	//	}
	//	temp := i
	//	defer f(temp)
	// }

	// 哪怕某个defer延迟调用发生错误，这些调用依旧会被执行

	// 滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里
	// testWithoutDefer()
	// testWithDefer()

	// defer陷阱
	// defer 与 closure
	// 如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值
	// _, _ = testDeferWithMulClosure(2, 0)

	// defer 与 return
	// testDeferWithReturn()

	// defer nil 函数
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	testDeferWithNil() // 名为 testDeferWithNil 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。然而值得注意的是，run() 的声明是没有问题，因为在testDeferWithNil函数运行完成后它才会被调用

	// 在错误的位置使用 defer
	useDeferIncorrectly()
	// 解决方案:   总是在一次成功的资源分配下面使用 defer ，对于这种情况来说意味着：当且仅当 http.Get 成功执行时才使用 defer
	useDeferCorrectly()

	// f.Close() 可能会返回一个错误，可这个错误会被我们忽略掉

	//

	//

}

// 一般使用
func usePanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型
		}
	}()

	panic("panic error!")

}

// 向已关闭的通道发送数据
func sendToClosedChannel() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
}

// 延迟调用中引发的错误
func errRaisedInDefer() {
	defer func() {
		fmt.Println(recover()) // 仅最后一个错误可被捕获
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}

func mulDeferRecover() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()
	defer recover()              // 无效！
	defer fmt.Println(recover()) // 无效！
	defer func() {
		func() {
			println("defer inner")
			recover() // 无效！
		}()
	}()

	panic("test panic")
}

func protectCodeSegmentWithAnonymousFunc(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

// 创建实现 error 接口的错误对象
var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}
func newErrorsAndTest() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}

func Try(fun func(), handler func(interface{})) { // 参数fun为执行函数，参数handler为异常处理函数
	defer func() {
		if err := recover(); err != nil {
			handler(err) // handler函数的参数为interface{}类型
		}
	}()
	fun()
}

// 类似 try catch 的异常处理
func tryCatch() {
	execFunc := func() { panic("test panic") }
	handleFunc := func(err interface{}) { fmt.Println(err) }
	Try(execFunc, handleFunc)
}

// 自定义异常
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}
func Open(filename string) error { // 返回值 error 为一个接口  类型PathError实现了error() string方法，即实现了这个接口

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

// panic & Recover
func panicRecoverBasics() {

	// panic
	//    1、内置函数
	//    2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
	//    3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
	//    4、直到goroutine整个退出，并报告错误

	// recover
	//	  1、内置函数
	//    2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
	//    3、一般的调用建议
	//        a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
	//        b). 可以获取通过panic传递的error

	// 注意：
	//     1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
	//     2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
	//     3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

	// 由于 panic、recover 参数类型为 interface{}，因此可抛出任何类型对象
	//    func panic(v interface{})
	//    func recover() interface{}
	usePanicRecover()

	// 向已关闭的通道发送数据会引发panic
	// sendToClosedChannel()

	// 延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获
	// errRaisedInDefer()

	// 捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递
	// mulDeferRecover()

	// 如果需要保护代码段，可将代码块重构成匿名函数，如此可确保后续代码被执
	// protectCodeSegmentWithAnonymousFunc(2, 1)

	// 除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。
	// 标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型
	// newErrorsAndTest()

	// Go实现类似 try catch 的异常处理
	tryCatch()

	// 如何区别使用 panic 和 error 两种方式?
	// 惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。

	// 自定义异常
	// 类型：PathError   定义error方法
	err := Open("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:

	}

}

// 单元测试
func unitTestBasics() {
	// 1.1. go test工具
	// go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件

	// 1.2. 测试函数
	// 格式：每个测试函数必须导入testing包，测试函数的名字必须以Test开头，参数类型必须为*testing.T，可选的后缀名必须以大写字母开头
	// 示例：定义一个split的包，包中定义了一个Split函数，再创建一个split_test.go的测试文件，并定义一个测试函数TestSplit(t *testing.T) {}， 包目录下执行 go test

	// go test命令添加-v参数，查看测试函数名称和运行时间
	// go test命令后添加-run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行

	// 注意：当我们修改了我们的代码之后不要仅仅执行那些失败的测试函数，我们应该完整的运行所有的测试，保证不会因为修改代码而引入了新的问题

	// 1.3. 测试组
	// 将待测函数的参数输入与返回值输出抽象为一个测试用例类型struct，多个测试即对应一个结构体数组，遍历操作这个结构体数组进行测试即可，如TestSplit2

	// 1.4. 子测试
	// 如果测试用例比较多的时候，我们是没办法一眼看出来具体是哪个测试用例失败，将上面的结构体数组改为map[string]struct{}
	// 同时Go1.7+中新增了子测试， 如：TestSplit3
	// 可以通过/来指定要运行的子测试用例，例如：go test -v -run=Split/simple只会运行simple对应的子测试用例

	// 1.5. 测试覆盖率
	// 测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例
	// Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率
	// -coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件

	// 1.6. 基准测试
	// 基准测试就是在一定的工作负载之下检测程序性能的一种方法
	// 1.6.1. 基准测试函数格式
	// 基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性
	// 1.6.2. 基准测试示例
	// 基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试，终端输出结果：其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值
	// -benchmem参数，来获得内存分配的统计数据
	// 1.6.3. 性能比较函数

	// 1.6.4. 重置时间

	// 1.6.5. 并行测试

	// 1.7. Setup与TearDown
	// 测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）

	// 1.7.2. 子测试的Setup与Teardown

	// 1.7.1. TestMain

	// 1.8. 示例函数

	// 1.8.1. 示例函数的格式

	// 1.8.2. 示例函数示例

}

// 压力测试
func pressureTestBasics() {

	// Go编写测试用例
	// 新建一个项目目录gotest,这样我们所有的代码和测试代码都在这个目录下，在该目录下面创建两个文件：gotest.go和gotest_test.go

	// 压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似
	// 压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
	// func BenchmarkXXX(b *testing.B) { ... }
	// go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数
	// 在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行 文件名也必须以_test.go结尾
}

func main() {

	// funBasics()
	// funcParamsWithUncertianLength("sum: %d", 1, 2, 3, 4, 5)
	// _slice := []int{1, 2, 3, 4, 5}
	// funcParamsWithUncertianLength("sum: %d", _slice...)    // 注意：使用 slice 对象做变参时，必须展开
	// funcParamsWithUncertianLength2("first params", 12, "third params", 32, [2]int{1, 2}, map[string]string{"name": "hu", "age": "22"})
	// anonymousFunc()
	// funcCalled()
	// dataTypeConvert()

	closureBasics()
	// recursiveFuncBasic()

	// deferBasics()

	// panicRecoverBasics()

	// unitTestBasics()

	// pressureTestBasics()
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
