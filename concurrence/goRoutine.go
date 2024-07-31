package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// 进程和线程
//		A. 进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
//		B. 线程是进程的一个执行实体,是CPU调度和分派的基本单位,它是比进程更小的能独立运行的基本单位。
//		C.一个进程可以创建和撤销多个线程;同一个进程中的多个线程之间可以并发执行。

// 多进程、多线程或者多协程开发，就少不了以下 3 个问题：
// 如何通信
// 如何访问共享资源
// 如何同步
// 多数编程语言多线程开发时需要用户写线程同步代码利用多个核，可能导致错误
// 而 Go 语言的 GMP 模型解决了多线程开发的问题，用户只关心协程的通信、共享、同步

// 并发和并行
//		A. 多线程程序在一个核的cpu上运行，就是并发。
//		B. 多线程程序在多个核的cpu上运行，就是并行。
// 协程和线程
//		协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
//		线程：一个线程上可以跑多个协程，协程是轻量级的线程。

// 并发主要由切换时间片来实现"同时"运行，并行则是直接利用多核实现多线程的运行，并行的"同时"是同一时刻可以多个进程在运行(处于running)，并发的"同时"是经过上下文快速切换，使得看上去多个进程同时都在运行而实际是交替运行的现象，是一种OS欺骗用户的现象
// 当程序中写下多进程或多线程代码时，这意味着的是并发而不是并行，并行与否程序员无法控制，只能让操作系统决定
// Go语言的并发模型是CSP，奉行通过通信来共享内存，而不是共享内存来通信

// 当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以
func hello1() {
	fmt.Println("sub goroutine running!")
}

// 启动单个goroutine
func run1() {
	// 程序启动时，main()函数作为主线程被创建------------------这里的线程严格的说指主goroutine，而非OS线程，下面亦然
	go hello1() // 创建子线程、执行子线程，这个过程主线程将继续往下执行
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second) // 若不稍等，主线程一结束，创建的协程有可能未执行完跟着销毁
}

var wg sync.WaitGroup // 实现goroutine的同步
func hello2(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

// 启动多个goroutine
func run2() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine，同步计数器的计数量+1
		go hello2(i)
	}
	wg.Wait() // 阻塞此主协程，等待所有登记的子goroutine结束后再继续执行
	fmt.Println("直到同步计数器计数量为0时才执行")
}

// 测试协程，主协程退出了，其他任务也退出
func testGoRoutine() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

// goroutine 与 OS线程
// 可增长的栈：
// 		OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）
// 		一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这个大
// goroutine调度：
// 		GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
// 			1.G就是goroutine，里面存放本goroutine信息，以及与所在P的绑定等信息。
// 			2.P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
// 			3.M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；
// 		P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。
// 		P的个数是通过runtime.GOMAXPROCS来设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。
// 		单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。

// Go语言中的操作系统线程和goroutine的关系：
// 		1.一个操作系统线程对应用户态多个goroutine。
// 		2.go程序可以同时使用多个操作系统线程。
// 		3.goroutine和OS线程是多对多的关系，即m:n

// runtime包的常用函数
// runtime.Gosched() 让出CPU时间片，重新等待安排任务
func testGosched() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched() // 生成处理器，让其他goroutine运行，但并不挂起当前goroutine
		fmt.Println("hello")
	}
}

// runtime.Goexit() 退出当前协程
func testGoexit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 终止当前协程，其他协程不受影响，且终止之前会执行所有的延迟调用defer
			runtime.Goexit() // return前打印输出了B、A两个defer
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	// 无限循环...
	for {
	}
}

// runtime.GOMAXPROCS
// Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上
// Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
// Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

// 将任务分配到不同的CPU逻辑核心上实现并行的效果
func testGOMAXPROCS(n int) {
	runtime.GOMAXPROCS(n)
	go a()
	go b()
	time.Sleep(time.Second)
}

// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
// 虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。
// Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
// 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

type Job struct {
	Id      int
	RandNum int
}
type Result struct {
	job *Job
	Sum int
}

// 自定义创建goroutine池函数
func createGoroutinePool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) { // 每个协程任务相同，都是从jobChan中取数据，进行计算后，放入resultChan
			for job := range jobChan {
				// 取jobChan中的随机数
				randNum := job.RandNum

				// 计算randNum各个位之和
				var sum int
				for randNum != 0 {
					sum += randNum % 10
					randNum /= 10
				}

				// 构造Result并传给resultChan
				resultChan <- &Result{
					job: job,
					Sum: sum,
				}
			}
		}(jobChan, resultChan)
	}
}

// Goroutine池
func goRoutinePoolDemo() {
	// 本质上是生产者消费者模型
	// 可以有效控制goroutine数量，防止暴涨

	// 需求：
	// 	计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
	// 	随机生成数字进行计算

	// 任务管道
	jobChan := make(chan *Job, 128)
	// 结果管道
	resultChan := make(chan *Result, 128)

	// 创建多个协程待命，当通道有数据时，这些协程会抢夺式获取并执行
	createGoroutinePool(64, jobChan, resultChan)

	// 开个打印结果的协程
	go func() {
		for res := range resultChan { // 阻塞，直到resultChan有值
			fmt.Println("jobID-----", res.job.Id, "\trandNum-----", res.job.RandNum, "\tSum-----", res.Sum)
		}
	}()

	var id int
	// 循环创建10000个job，输入到jobChan
	for id < 10000 {
		id++
		job := &Job{
			Id:      id,
			RandNum: rand.Int(),
		}
		jobChan <- job
	}

}

// 定时器
func timerAndTickerBasics() {
	// Timer：时间到了，只执行1次
	// 1.timer基本使用
	// timer1 := time.NewTimer(2 * time.Second)
	// t1 := time.Now()
	// fmt.Println("t1-------", t1)
	// timer.C  类型为 <-chan Time  即一个单向通道
	// t2 := <- timer1.C                     // 从timer.C这个通道中取值，即是定时sleep的时候
	// fmt.Printf("t2-------%v", t2)

	// 2.验证timer只能响应1次
	// timer2 := time.NewTimer(2 * time.Second)
	// defer timer2.Stop()                   // 在这里并不会停止定时器。这是因为Stop会停止Timer，停止后，Timer不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误，如果想停止定时器，只能让go程序自动结束
	// for {
	// 	<- timer2.C                          // 定时2秒
	// 	timer2.Reset(time.Second)            // 重置计数时间，定时器设置定时1秒
	// 	fmt.Println("时间到")                 // 当不Reset时仅打印一次即响应一次，后面的循环就报错deadlock，原因是通道阻塞
	// }

	// 3.timer实现延时的功能
	// time.Sleep(time.Second)
	//
	// timer3 := time.NewTimer(2 * time.Second)
	// t3 := <-timer3.C
	// fmt.Println("2秒到", t3)
	//
	// af := <-time.After(2 * time.Second)
	// fmt.Println("2秒到", af)

	// 4.停止定时器
	// timer4 := time.NewTimer(2 * time.Second)
	// go func() {
	// 	<-timer4.C
	// 	fmt.Println("定时器执行了")
	// }()
	// timer.Stop() 阻止timer被触发，成功返回true，若timer已过期或已被stop过，返回false。
	// 注意：Stop仅阻止定时器触发，而不会关闭通道，以保证通道数据的正常读取
	// if ok := timer4.Stop(); ok {
	// 	fmt.Println("timer4已经关闭")
	// }

	// 5.重置定时器
	// timer5 := time.NewTimer(5 * time.Second)
	// timer5.Reset(2 * time.Second)                // 将上面5秒的定时器重置为2秒的定时器
	// fmt.Println(time.Now())
	// fmt.Println(<-timer5.C)

	// Ticker：每隔一定时间就会执行
	// ticker := time.NewTicker(1 * time.Second)
	// var i int
	// go func() {
	// 	for {
	// 		i ++
	// 		fmt.Println(<-ticker.C)
	// 		if i == 5 {
	// 			// 停止
	// 			ticker.Stop()
	// 		}
	// 	}
	// }()
	// for {        // 保持主协程alive
	// }

	// time.After直接返回一个定时器的通道
	ch := time.After(2 * time.Second)
	af := <-ch // 表示2秒后返回一条time.Time类型的通道消息
	fmt.Println("2秒到", af)

}

// select多路复用
func selectMultiplexing() {
	// 某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞

	// select关键字，可以同时响应多个通道的操作，select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句
	// select可以同时监听一个或多个channel，直到其中一个channel ready
	ch1 := make(chan string, 5)
	ch2 := make(chan string, 5)

	// 创建协程函数
	routine := func(ch chan<- string, data string, sep time.Duration) {
		fmt.Println("params------", data, sep)
		time.Sleep(sep)
		ch <- data
	}
	// 跑2个子协程，写数据
	go routine(ch1, "data1", 2*time.Second)
	go routine(ch2, "data2", 1*time.Second)

	// 用select监控，ch1/ch2哪个通道先取到数据就执行对应的case
	var counter int // 用于确定s1，s2两个数据是否都取到
	var s1 string
	var s2 string
	for {
		select {
		case s1 = <-ch1: // 从ch1取值
			counter++
		case s2 = <-ch2: // 从ch2取值
			counter++
		}
		if counter == 2 {
			break
		}
	}
	fmt.Println(s1, s2)

}

// 数据竞态--------------多个goroutine同时操作一个资源（临界区）
func testDataRace() {
	var x int64
	var wg sync.WaitGroup

	add := func() {
		for i := 0; i < 5000; i++ {
			x = x + 1
		}
		wg.Done()
	}

	wg.Add(2)
	go add()
	go add()
	wg.Wait() // 等待两个add协程执行完毕才继续执行
	fmt.Println(x)
}

// 互斥锁------------------------解决上述竞态问题
func useMutex() {

	var x int64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	add := func() {
		for i := 0; i < 5000; i++ {
			mutex.Lock() // 加锁
			x = x + 1
			mutex.Unlock() // 解锁
		}
		wg.Done()
	}
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

// 读写互斥锁--------------------适合读多写少的场景
func useRWMutex() {
	var (
		x  int64
		wg sync.WaitGroup
		// lock   sync.Mutex
		rwMutex sync.RWMutex
	)
	write := func() {
		// lock.Lock()                                 // 加互斥锁
		rwMutex.Lock() // 加写锁
		x = x + 1
		time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
		rwMutex.Unlock()                  // 解写锁
		// lock.Unlock()                               // 解互斥锁
		wg.Done()
	}
	read := func() {
		// lock.Lock()                                 // 加互斥锁
		rwMutex.RLock()              // 加读锁
		time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
		rwMutex.RUnlock()            // 解读锁
		// lock.Unlock()                               // 解互斥锁
		wg.Done()
	}

	startTime := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	endTime := time.Now()
	fmt.Println("endTime - startTime = ", endTime.Sub(startTime))

}

// 比较互斥锁和原子操作的性能
func compareMutexWithAtomic() {
	var x int64
	var wg sync.WaitGroup
	// var m sync.Mutex

	// 普通版加函数
	// add := func() {
	// 	x++
	// 	wg.Done()
	// }
	// 互斥锁版加函数
	// mutexAdd := func() {
	// 	m.Lock()
	// 	x++
	// 	m.Unlock()
	// 	wg.Done()
	// }
	// 原子操作版加函数
	atomicAdd := func() {
		atomic.AddInt64(&x, 1)
		wg.Done()
	}

	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		// go add()              // 普通的add函数 不是并发安全的，结果非预期，具有不确定性
		// go mutexAdd()         // 加互斥锁的add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 加原子操作的add函数 是并发安全，性能优于互斥锁
	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println(x)
	fmt.Println(endTime.Sub(startTime))

}

// 并发安全和锁问题
func concurrentSecurityAndLock() {
	// 有时候会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题，举个例子：
	// testDataRace()                 // 开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致结果与期待的10000不符

	// 1.1.1. 互斥锁
	// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。
	// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的
	// 使用互斥锁来修复上面testDataRace()代码的问题：
	useMutex()

	// 1.1.2. 读写互斥锁
	// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择
	// 读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
	// 读写锁分为两种：读锁和写锁。
	// 		当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
	// 		当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待
	// useRWMutex()

	// 原子操作
	// 加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高
	// 可以使用原子操作来保证并发安全，它在用户态就可以完成，因此性能比加锁操作更好
	// compareMutexWithAtomic()

}

// 使用 Sync.WaitGroup 实现协程的同步
func useSyncWaitGroup() {

	// sync.WaitGroup也是一个经常会用到的同步方法，它的使用场景是在一个goroutine等待一组goroutine执行完成
	// sync.WaitGroup拥有一个内部计数器。当计数器等于0时，则Wait()方法会立即返回。否则它将阻塞执行Wait()方法的goroutine，直到计数器等于0时为止。

	var wg sync.WaitGroup
	wg.Add(1) // 增加计数器，将计数器加2，因为开了两个add协程，每个协程为一个计数量
	go func() {
		fmt.Println("goroutine is executing ...")
		wg.Done() // 减少计数器，此add方法执行完成，将计数器减1，Done()方法底层就是通过Add(-1)实现的
	}()
	fmt.Println("main goroutine wait until sync Counter become zero ...")
	wg.Wait() // 阻塞当前调用wait()的goroutine，直到计数器等于0时为止，否则当前goroutine结束了，其子goroutine也被杀死
	fmt.Println("main goroutine done!")
}

// 并发安全问题
func testConcurrentSecurity() {
	var once sync.Once
	var _map map[string]interface{}

	// 初始化操作函数
	initOperationFunc := func() {
		_map = map[string]interface{}{
			"left":  "left111",
			"up":    123,
			"right": "right222",
			"down":  447,
		}
	}

	// concurrenceFunc异步函数   被多个 goroutine 调用时不是并发安全的
	concurrenceFunc := func() {
		if _map == nil {
			// initOperationFunc()       // 实际操作数据之前先判断数据是否已初始化，没有则进行初始化操作

			// sync.Once解决testConcurrentSecurity函数遇到的并发安全问题
			once.Do(initOperationFunc)
		}

		// 操作_map数据
		fmt.Println("--------- operate initialized data ---------")
		for k, v := range _map {
			fmt.Println(k, v)
		}
	}

	// 多个concurrenceFunc异步执行
	for i := 0; i < 5; i++ {
		go concurrenceFunc() // 多个goroutine并发调用concurrenceFunc函数时不是并发安全的，现代的编译器和CPU可能会在保证每个goroutine都满足串行一致的基础上自由地重排访问内存的顺序
		// 有可能出现：某个goroutine执行初始化操作执行到一半，_map非nil但未初始化完成，而此时另一个goroutine判断_map非空继续往下执行，拿着非空但未初始化完整的_map进行后续操作，这就出现了并发安全问题。
		// 考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化_map的时候不会被其他的goroutine操作，但是这样做又会引发性能问题，于是sync.Once.Do(initOperationFunc)来解决
	}

	for {

	}
}

// 测试go内置的map不是并发安全的
func testBuiltinMap() {
	wg := sync.WaitGroup{}

	var m = make(map[string]int)
	get := func(key string) int {
		return m[key]
	}
	set := func(key string, value int) {
		m[key] = value
	}

	func(num int) { // num表示开启的goroutine数，开启少量没有问题，开启大量则报错 fatal error: concurrent map writes
		for i := 0; i < num; i++ {
			wg.Add(1)
			go func(n int) {
				key := strconv.Itoa(n) // 将整型转为整型的字符串： 1 -> '1'
				set(key, n)
				fmt.Printf("k=:%v,v:=%v\n", key, get(key))
				wg.Done()
			}(i)
		}
		wg.Wait()
	}(2)

}

// 大量并发操作map时，使用sync.Map代替内置map
func useSyncMap() {
	wg := sync.WaitGroup{}
	var syncMap = sync.Map{}

	func(num int) {
		for i := 0; i < num; i++ {
			wg.Add(1)
			go func(n int) {
				key := strconv.Itoa(n)
				syncMap.Store(key, n)         // 相当于根据key设置value
				value, _ := syncMap.Load(key) // 相当于根据key获取value
				fmt.Printf("k=:%v,v:=%v\n", key, value)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}(20)

}

// 并发中的同步
func syncInConcurrence() {
	// 1.1.1. sync.WaitGroup
	// 使用sync.WaitGroup来实现并发任务的同步
	// sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。
	// 例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。
	// var wg sync.WaitGroup        // 注意sync.WaitGroup是一个结构体，传递的时候要传递指针
	// wg.Add(2)                    // 计数器+delta
	// wg.Done()                    // 计数器-1
	// wg.Wait()                    // 阻塞当前调用wait的协程，直到计数器变为0才继续执行
	useSyncWaitGroup()

	// 1.1.2. sync.Once
	// 在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等
	// sync包中提供了一个针对只执行一次场景的解决方案 – sync.Once
	// sync.Once 内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次
	// sync.Once只有一个Do方法，声明如下
	// func (o *Once) Do(f func()) {}          注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用
	// var once sync.Once
	// once.Do(func() {})

	// 加载配置文件示例
	// 对于一个开销很大的初始化操作，一开始不执行初始化，而是延迟执行(真正用到它的时候再执行)，这是比较合理的，原因是预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，还有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的，例如：
	// testConcurrentSecurity()      // 注释once.Do(initOperationFunc)

	// 使用sync.Once 解决上述testConcurrentSecurity函数遇到的并发安全问题
	// testConcurrentSecurity()      // 使用once.Do(initOperationFunc)

	// 1.1.3. sync.Map
	// Go语言中内置的map不是并发安全的
	// testBuiltinMap()

	// 上述testBuiltinMap场景下，就需要为内置map加锁来保证并发的安全性了
	// Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法
	// useSyncMap()

}
func main() {
	// run1()
	// run2()

	// testGoRoutine()

	// runtime包的使用
	// testGosched()
	// testGoexit()
	// testGOMAXPROCS(1)                     // 两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	// testGOMAXPROCS(2)                     // 将逻辑核心数设为2，此时两个任务并行执行

	// goRoutinePoolDemo()

	// timerAndTickerBasics()

	// selectMultiplexing()

	concurrentSecurityAndLock()

	// syncInConcurrence()

}
