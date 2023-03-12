package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	fmt.Println("======= readFile =============")
	const filename = "../assets/abc.txt"

	// contents, err := ioutil.ReadFile(filename)
	// // if 判断不需要（）
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	// 优化 （if根多个语句）
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// contents 在if作用域里面 外面无法访问
	// fmt.Printf("%s\n", contents)
}

// switch 后面可以不跟表达式
func grade(score int) string {
	fmt.Println("======== grade =============")
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}
