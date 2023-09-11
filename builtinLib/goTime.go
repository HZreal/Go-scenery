package main

import (
	"fmt"
	"time"
)

func main() {
	// UTC时间1970年1月1日0时0分45秒,所以会打印出 t:45
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "1970-01-01 00:00:45", time.UTC)
	fmt.Println("t:", t.Unix())

	// 3. 秒数转Duration
	t1 := time.Now()                             // 当前时间
	time.Sleep(3 * time.Second)                  // 休眠3秒
	dur1 := 2*time.Second + 500*time.Millisecond // 规定间隔:2秒500毫秒
	// 判断间隔时间
	if time.Now().Sub(t1) > dur1 {
		fmt.Println("sleep over 2.5s")
	}

	// 4. time.Duration 转秒数
	dur2 := 2*time.Second + 500*time.Millisecond
	durSec := dur2.Seconds()
	fmt.Println("durSec:", durSec)

}
