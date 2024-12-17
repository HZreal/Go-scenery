package main

/**
 * @Author HZreal
 * @Date 2024-12-17
 * @File: interface.go
 * @Description: 参考：https://github.com/xxjwxc/uber_go_guide_cn
 */

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

// 不要使用指向接口类型的指针，接口本身就是引用类型，接口底层还是指针
// 若需要接口方法修改基础数据，则必须使用指针传递 (将对象指针赋值给接口变量)

// ---------------------------------------------------------------------------------------------------------------

// 编译阶段验证接口的实现类型的合法性，避免运行时检查错误!!!

type MyHandler struct {
	// ...
}

// 用于触发编译期的接口的合理性检查机制
// 如果 MyHandler 没有实现 http.Handler，会在编译期报错
// 对于指针类型实现（如 *Handler、切片和映射），赋值的右边应该是断言类型的零值。
var _ http.Handler = (*MyHandler)(nil)

func (h *MyHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

type LogHandler struct {
	h   http.Handler
	log *zap.Logger
}

// 对于非指针类型实现，赋值的右边应该是空值。
var _ http.Handler = LogHandler{}

func (h LogHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}

// ---------------------------------------------------------------------------------------------------------------

// 使用值接收器的方法既可以通过值调用，也可以通过指针调用
// 带指针接收器的方法只能通过指针或 addressable values 调用

type S struct {
	data string
}

// 使用类型值实现的，类型值和类型指针均可调用
func (s S) Read() string {
	return s.data
}

// 使用类型指针实现的，只有类型指针才可以调用
func (s *S) Write(str string) {
	s.data = str
}

func demo1() {
	// 创建结构体值类型
	sVals := map[int]S{1: {"A"}}

	// 通过类型值调用 Read，无问题
	sVals[1].Read()

	// 这不能编译通过，因为 Write 是类型指针实现的，而 sVals[1] 是类型值
	// sVals[1].Write("test")

	// 为什么能通过？
	// 带指针接收器的方法除了能通过类型指针调用，还可以通过可寻址的值（addressable values）调用
	// 参考 https://go.dev/ref/spec#Method_values
	// 而 a 是可寻址的值，编译器隐式的找到了其指针
	a := sVals[1]
	a.Write("test")

	// bPtr := &sVals[1] // 报错原因是 map 的值是非指针，无法直接对索引取址，详见 map 中有说明
	// bPtr.Write("test")
	b := sVals[1] // 进行了赋值就可以取址了，但注意是值拷贝
	bPtr := &b
	bPtr.Write("test")

	// 创建结构体指针类型
	sPtrs := map[int]*S{1: {"A"}}

	// 通过类型指针既可以调用 Read，也可以调用 Write 方法
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}

// ---------------------------------------------------------------------------------------------------------------

type F interface {
	f()
}

type S1 struct{}

// S1 使用类型值实现
func (s S1) f() {}

type S2 struct{}

// S2 使用类型指针实现
func (s *S2) f() {}

func demo2() {
	s1Val := S1{}
	s1Ptr := &S1{}
	// s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr

	//  下面代码无法通过编译。因为 s2Val 是一个值，而 S2 的 f 方法中没有使用值接收器
	// i = s2Val

	fmt.Println(i)
}

func main() {
	demo1()
	demo2()
}
