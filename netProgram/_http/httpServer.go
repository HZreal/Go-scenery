package main

import (
	"fmt"
	"github.com/kardianos/osext"
	"net/http"
)

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println("method ---->  ", r.Method)
	fmt.Println("url ---->  ", r.URL.Path)
	fmt.Println("header ---->  ", r.Header)
	fmt.Println("body ---->  ", r.Body)

	// 响应
	w.Write([]byte("handler response !"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 获取执行二进制文件所在的目录
	folderPath, _ := osext.ExecutableFolder()
	fmt.Println("本go文件编译成的二进制文件的执行目录 ---->  ", folderPath)
	http.ServeFile(w, r, folderPath+"/index.html")
}

func main() {

	// 回调函数
	http.HandleFunc("/go", myHandler) // http://127.0.0.1:8000/go
	http.HandleFunc("/", indexHandler)

	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}
