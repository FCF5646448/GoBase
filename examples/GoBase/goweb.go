package main

import (
	"net/http"
	"log"
	"fmt"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintf(w,"hello astaxie")
}

func main() {
	http.HandleFunc("/",sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil  {
		log.Fatal("ListenAndServer:",err)
	}
}
