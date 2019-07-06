package main

import "net/http"

func main(){
	// 读取当前目录文件
	http.ListenAndServe(":12345",http.FileServer(http.Dir(".")))
}
