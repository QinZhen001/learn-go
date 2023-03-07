package main

import "fmt"

// () 集中定义
var (
	aaa = 3
	kkk = "kkk"
)

// 这个作用域不能使用 :=  (包内部变量 没有全局变量的说法)
// bbb := "a"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

func variableTypeDeduction() {
	var a, b, c = 3, "aaa", false
	fmt.Println(a, b, c)
}

func variableShorter() {
	// 只有初值才能 :=
	a, b, c := 3, "aaa", false
	a = 5
	fmt.Println(a, b, c)
}
