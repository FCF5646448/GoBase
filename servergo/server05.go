/*  实验返回一个xml对象，对象里包含字典
 */
package main

import (
	"net/http"
	"encoding/xml"
)

type XmlData struct {
	Name string
	Hobbies []string
}


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/xml",xmlHandle)

	http.ListenAndServe(":12345",mux)
}

func xmlHandle(w http.ResponseWriter, r *http.Request) {
	data := XmlData{"fcf",[]string{"1","2","3"}}
	x, err := xml.MarshalIndent(data,""," ")
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/xml")
	w.Write(x)
}
