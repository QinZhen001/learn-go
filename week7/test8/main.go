package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	ch := make(chan int)

	select {
	case <-ch:
		fmt.Println("read from ch", <-ch)
	case <-timer.C:
		println("3 seconds")
	}

}
