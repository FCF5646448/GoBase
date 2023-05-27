package main

import "net/http"

func main() {
	// r 代表来自客户端的请求
	// w 用于操作返回给客户端的应答
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})
	// 监听本地8080端口
	http.ListenAndServe(":8080", nil)
}
