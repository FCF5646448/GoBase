/*  实验返回一个json对象，对象里包含字典
 */
package main

import (
	"net/http"
	"encoding/json"
)

type Data struct {
	Name string 	`json:"name"`
	Hobbies []string `json:"hobbies"`
}

func listHandle(w http.ResponseWriter, r *http.Request) {
	data := Data{"fcf",[]string{"吃饭","睡觉"}}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type","application/json")
	w.Write(js)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/list",listHandle)

	http.ListenAndServe(":12345",mux)
}



