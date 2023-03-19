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
	fmt.Println("========= writeFile ============")
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

// create err
func writeFileErr(filename string) {
	fmt.Println("========= writeFileErr ============")
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println("========= pathError ============")
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	defer file.Close()
}

func main() {
	tryDefer()
	filename := "./fib.txt"
	writeFile(filename)
	writeFileErr(filename)
}
