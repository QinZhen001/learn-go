package main

import "fmt"

type CheckOut func(int, int) int

func Show() {
	fmt.Println("show")
}

func main() {
	// 函数变量
	Show()
	show := Show
	show()

	// 匿名函数
	show2 := func() {
		fmt.Println("show2")
	}
	show2()

	var checkOut CheckOut = func(a int, b int) int {
		return a + b
	}
	fmt.Println(checkOut(68, 98))

}
