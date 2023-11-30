package main

import "fmt"

func main() {
	// some code ...

	// try to recover from panic
	defer func() {
		fmt.Println("try to recover from panic")
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	panic("a problem")
}
