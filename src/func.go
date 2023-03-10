package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// 可以返回多个值
// (int, int) 也可以起别名 (q, r int)
func div(a int, b int) (q, r int) {
	fmt.Println("============ div ================")
	// return a / b, a % b
	q = a / b
	r = a % b
	return // 当函数比较长的时候不建议这种写法
}

// switch 会自动break
func eval(a, b int, op string) (int, error) {
	fmt.Println("======== eval =============")
	var result int
	var error error
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return 0, fmt.Errorf("unsupported operator: " + op)
	}

	return result, error
}

func add(a int, b int) int {
	return a + b
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Println("============ op =================")
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	fmt.Println("=========== sum ==================")
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
