package main

import (
	"bufio"
	"fmt"
	"os"
)

// defer
// 确保调用在函数结束时发生
// 参数在defer语句时计算
// defer 列表为先进后出

func tryDefer() {
	fmt.Println("========= tryDefer ============")
	defer fmt.Println(2)
	fmt.Println(1)
	defer fmt.Println(3)
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := Fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("./fib.txt")
}
