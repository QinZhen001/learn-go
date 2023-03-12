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
	v1, err1 := eval(5, 3, "*")
	fmt.Println(v1, err1)
	v2, err2 := eval(5, 3, "k")
	fmt.Println(v2, err2)
	if err2 != nil {
		fmt.Println("eval error")
	}
	g := grade(88)
	fmt.Println(g)

	// loop
	fmt.Println(convertToBin(13))
	printFile1("../assets/abc.txt")

	// func
	q, _ := div(10, 3) // _ 代表不想要这个值
	fmt.Println(q)
	apply(add, 1, 2)
	fmt.Println(sum(1, 2, 4))

	// pointer
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	c, d := 3, 4
	c, d = swap1(c, d)
	fmt.Println(c, d)

	// arr
	test_arr1()

	// slice
	slice()
	create_slice()

	// map
	create_map()

	// string
	create_string()

}
