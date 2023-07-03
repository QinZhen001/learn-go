package main

import (
	"errors"
	"fmt"
)

func funcRecover() error {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Panic recover! v:%v", v) // 3
		}
	}()

	return funcCook()
}

func funcCook() error {
	panic("funcCook() panic!") // 2
	return errors.New("发生错误了")
}

func main() {
	err := funcRecover()
	if err != nil {
		fmt.Println(err)
	}

}
