/*  实验返回文件, 需在当前路径下，建一个images文件夹，然后放入11.jpeg 图片
 */
package main

import (
	"net/http"
	"path"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/file",fileHandle)

	http.ListenAndServe(":12345",mux)
}

func fileHandle(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("images","11.jpeg")
	http.ServeFile(w,r,fp)
}