package main

import (
	"context"
	"fmt"
	"time"
)

// timer 某个时刻触发的定时器
func timerUsage() {
	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	fmt.Println("timer expired after 5 seconds")
}

func timerStop() {

}

func timerReset() {

}

func timeAfterFunc() {
	f := func() {
		fmt.Println("do while f is invoked")
	}

	timer := time.AfterFunc(time.Second*2, f)

	defer timer.Stop()

	time.Sleep(time.Second * 4)
}

func timeAfter() {

}

// ///////////////////////// ticker  //////////////////////
// 固定间隔触发的定时器
func tickerUsage() {
	ticker := time.NewTicker(time.Second * 2)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, ticker *time.Ticker) {
		defer ticker.Stop() // 注意要 stop ，否则容易造成内存泄露
		for {
			select {
			case <-ticker.C:
				fmt.Println("exec every 2 seconds")
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			}
		}
	}(ctx, ticker)

	time.Sleep(time.Second * 10)
	cancel()
}

func main() {
	// timerUsage()
	// tickerUsage()
	// timeAfterFunc()
	// timeAfter()
	tickerUsage()

}
