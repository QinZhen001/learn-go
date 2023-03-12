package main

import (
	"fmt"
	"learngo/interface/mock"
	"learngo/interface/real"
)

var url = "https://www.baidu.com"

// duck typing
// 像鸭子走路，长得像鸭子，就是鸭子
// 描述事物的外部行为而非内部结构

// go 属于结构化类型系统

// 接口由使用者定义
// 接口变量自带指针
// 接口变量同样采用值传递，几乎不需要使用接口的指针
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(p Poster) string {
	return p.Post(url, map[string]string{
		"name": "qdg",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(r RetrieverPoster) string {
	r.Post(url, map[string]string{
		"contents": "new contents",
	})

	return r.Get(url)
}

// 任何类型
type Queue []interface{}

// 如果返回值要求 int
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	// 强制转成 int
	return head.(int)
}

func main() {
	retriever := mock.Retriever{Contents: "mock text"}
	r1 := &retriever
	fmt.Printf("retriever=  %T %v \n", retriever, retriever)
	fmt.Printf("r=  %T %v \n", r1, r1)
	fmt.Println("download= ", download(r1))
	fmt.Println()
	// var r2 Retriever
	// r2 = mock.Retriever{Contents: "mock text"}
	// fmt.Println(r2.type)
	// switch r2.(type) {

	// }
	fmt.Println()
	r3 := real.Retriever{}
	fmt.Println(download(r3))
	fmt.Println()
	fmt.Println("========== type conversion ===========")
	q := Queue{1, 2, 3}
	fmt.Println(q.Pop())
	fmt.Println("========== get and post ==================")
	fmt.Println(session(r1))
}
