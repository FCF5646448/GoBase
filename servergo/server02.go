/*  实验实验SeverMux 路由
 */
package main

import "net/http"

func helloHandler(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("hello world"))
}

func main(){
	http.HandleFunc("/",helloHandler)
	http.ListenAndServe(":12345",nil)
}
