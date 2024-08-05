package main

/**
 * @Author elastic·H
 * @Date 2024-08-04
 * @File: serverCase.go
 * @Description: Server 情况的分析
 */

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // pprof 对 Server 型的分析时，只需要引入加载即可
)

func handler(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(resp, "Hello World")
}

func demo() {
	http.HandleFunc("/", handler)
	// 注意：ListenAndServe 的第二个参数为 handler
	// 若为 nil，会自动注册 pprof 处理路由； 若不为 nil，则需要手动注册
	log.Fatal(http.ListenAndServe(":28080", nil))
	// 通过 UI 进行分析：浏览器输入 http://localhost:28080/debug/pprof/

	// 若需要手动注册这几个函数
	// http.HandleFunc("/debug/pprof/", pprof.Index)
	// http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// http.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

// RegisterPprof 通过协程另起一个HTTP服务，单独用作pprof分析
func RegisterPprof() {
	go func() {
		if err := http.ListenAndServe(":28080", nil); err != nil {
			panic("pprof server start error: " + err.Error())
		}
	}()
}

// 在生产环境中通常通过协程另起一个服务用于 pprof 分析
func useInProdEnv() {
	RegisterPprof()

	for {
	}
}

func main() {
	// demo()
	useInProdEnv()
}
