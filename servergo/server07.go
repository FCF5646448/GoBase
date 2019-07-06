/*  实验返回html, 需在当前路径下，建一个templates文件夹，然后新建index.html 失败
 */
package main

import (
	"net/http"
	"path"
	"html/template"
)

type DataHtml struct {
	Name string
	Hobbies []string
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/html",htmlHandle)

	http.ListenAndServe(":123456",mux)
}

func htmlHandle(w http.ResponseWriter, r *http.Request) {
	hl := DataHtml{"fcf",[]string{"1","2"}}

	fp := path.Join("templates","index.html")
	tmp,err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	if err := tmp.Execute(w,hl); err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
}
