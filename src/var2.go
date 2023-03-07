package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Println("===== euler ====")
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	// (0+1.2246467991473515e-16i)  并不是0
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	// 更好的写法  并不是0  (浮点数不准)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	// 保留3位小数
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// go 只能强制类型转换
func triangle() {
	fmt.Println("======= triangle ===============")
	var a, b, c int = 3, 4, 0
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 定义常量  (常量不用大写定义 在go中首字母大写是有特殊含义的)
func consts() {
	fmt.Println("============ const ==============")
	const filename string = "abc.text"
	// 不规定类型 就是不确定类型
	var c int
	const a, b = 3, 4 // 变为 float
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
}

// 枚举值 （一般用 const）
func enums() {
	fmt.Println("============== enums ===============")
	// const (
	// 	cpp  = 0
	// 	java = 1
	// 	js   = 2
	// 	py   = 3
	// )
	const (
		cpp = iota
		_   // 跳过一个值
		java
		js
		py
	)
	fmt.Println(cpp, java, js, py)
}
