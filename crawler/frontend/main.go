package main

import (
	"learngo/crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
