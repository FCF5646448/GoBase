package main

import (
	"net/http"
	"fmt"
)

func main(){
	//第一个参数是接口，
	http.HandleFunc("/",helloworld)
	// 这里默认是127.0.0.1
	http.ListenAndServe(":8081",nil)
}

func helloworld(rw http.ResponseWriter, req *http.Request){



	//返回字符串
	fmt.Fprint(rw, "hello world")
}