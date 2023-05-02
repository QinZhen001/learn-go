package main

import (
	"learngo/crawler/frontend/controller"
	"net/http"
)

func main() {
	// 文件服务器
	http.Handle("/", http.FileServer(http.Dir("view")))
	http.Handle("/search", controller.CreateSearchResultHandler("view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
