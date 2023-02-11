package main

import (
	"net/http"
	"fmt"
)


func main(){
	mux := http.NewServeMux()
	//mux.HandleFunc("/hi",sayHi)

	mux.Handle("/hi2", http.HandlerFunc(sayHi))

	http.ListenAndServe(":9000",mux)
}

func sayHi(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("content-type","application/json")
	//返回字符串
	response.Write([]byte("ni hao"))
	fmt.Fprintf(response,"ni hao")
}
