package main

import (
	"context"
	"fmt"
	"time"
)

// select + chan 控制 goroutine
func controlGoroutineBySelectAndChannel() {
	// 主协程通过往 channel 发送开关，从而控制子协程的退出

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("goroutine exit")
				return
			default:
				fmt.Println("监控中 ...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Println("准备通知监控停止 ...")
	stop <- true
	time.Sleep(time.Second * 5)
}

// context 控制协程优雅退出、传递上下文信息

// type Context interface {
//	Deadline() (deadline time.Time, ok bool)
//
//	Done() <-chan struct{}
//
//	Err() error
//
//	Value(key interface{}) interface{}
// }

// 通过 context 改写上述控制协程的退出
func controlGoroutineByContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine exit")
				return
			default:
				fmt.Println("监控中 ...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(time.Second * 10)

	fmt.Println("准备通知监控停止 ...")
	cancel()

	time.Sleep(time.Second * 5)
}

// context 控制多个协程
func controlMultiGoroutineByContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "监控子协程1")
	go watch(ctx, "监控子协程2")
	go watch(ctx, "监控子协程3")

	time.Sleep(10 * time.Second)

	fmt.Println("准备通知监控停止 ...")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

// Context 的继承衍生
// WithCancel 获取取消函数
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// WithDeadline 设置过期截止时间（时间点）
// func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
// WithTimeout 设置多久后过期（时间段）
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// WithValue 传递必要的上下文参数（参数要是线程安全的）
// func WithValue(parent Context, key, val interface{}) Context

func main() {
	// controlGoroutineBySelectAndChannel()
	// controlGoroutineByContext()
	controlMultiGoroutineByContext()
}
