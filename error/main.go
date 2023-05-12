package main

import (
	"errors"
	"fmt"
)

func funcRecover() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()

	return funcCook()

	// panic(errors.New("this is an error"))
}

func funcCook() error {
	panic("停水了")
	return errors.New("this is an error") // 不执行
}

// 错误(业务) 和 异常
func main() {
	// panic recover defer
	err := funcRecover()
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}
}
