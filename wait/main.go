package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.WaitGroup Add Done Wait
	s := []int{1, 2, 3}
	var wg sync.WaitGroup
	for _, item := range s {
		wg.Add(1)
		go SayHello(item, &wg)
	}
	wg.Wait()
	fmt.Println("main end")
}

func SayHello(item int, wg *sync.WaitGroup) {
	fmt.Println("hello", item)
	wg.Done()
}
