package main

/**
 * @Author elastic·H
 * @Date 2024-07-31
 * @File: trace.go
 * @Description:
 */

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}

// 编译 go build trace.go
// 运行 GODEBUG=schedtrace=1000 ./trace

// 输出如下：
// GODEBUG=schedtrace=1000 ./trace
// SCHED 0ms: gomaxprocs=10 idleprocs=9 threads=2 spinningthreads=0 needspinning=0 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0 0 0]
// Hello World
// SCHED 1006ms: gomaxprocs=10 idleprocs=10 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0 0 0]
// Hello World
// SCHED 2007ms: gomaxprocs=10 idleprocs=10 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0 0 0]
// Hello World
// SCHED 3013ms: gomaxprocs=10 idleprocs=10 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0 0 0]
// Hello World
// SCHED 4015ms: gomaxprocs=10 idleprocs=10 threads=5 spinningthreads=0 needspinning=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0 0 0]
// Hello World

// 解释
// SCHED：调试信息输出标志字符串，代表本行是 goroutine 调度器的输出；
// 0ms：即从程序启动到输出这行日志的时间；
// gomaxprocs: P 的数量，本例有 10 个 P, 因为默认的 P 的属性是和 cpu 核心数量默认一致，当然也可以通过 GOMAXPROCS 来设置；
// idleprocs: 处于 idle 状态的 P 的数量；通过 gomaxprocs 和 idleprocs 的差值，我们就可知道执行 go 代码的 P 的数量；
// threads: os threads/M 的数量，包含 scheduler 使用的 m 数量，加上 runtime 自用的类似 sysmon 这样的 thread 的数量；
// spinningthreads: 处于自旋状态的 os thread 数量；
// idlethread: 处于 idle 状态的 os thread 的数量；
// runqueue=0： Scheduler 全局队列中 G 的数量；
// [0 0 0 0 0 0 0 0 0 0]: 分别为 10 个 P 的 local queue 中的 G 的数量。
