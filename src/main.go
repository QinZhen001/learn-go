package main

import "fmt"

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aaa, kkk)

	// var2
	euler()
	triangle()
	consts()
	enums()

	// var3
	readFile()
	r := eval(5, 3, "*")
	fmt.Println(r)
	g := grade(88)
	fmt.Println(g)
}
