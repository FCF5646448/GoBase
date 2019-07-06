/*  实验实验SeverMux 路由
 */
package main

import (
	"net/http"
	"io"
)



func helloHanlde(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"hello")
}

func hiHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main(){
	//路由
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",helloHanlde)
	mux.HandleFunc("/hi",hiHandle)

	http.ListenAndServe(":12345",mux)
}