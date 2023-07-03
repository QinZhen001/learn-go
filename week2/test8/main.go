package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []string{"a", "b", "c"}
	var wg sync.WaitGroup
	for _, item := range s {
		wg.Add(1)
		go func(item string) {
			println(item)
			wg.Done()
		}(item)
	}
	wg.Wait()
	fmt.Println("您点的菜上齐了")

}
