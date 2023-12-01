package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				return
			default:
				fmt.Printf("message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	// has closed channel
	time.Sleep(3 * time.Second)
	fmt.Println("main process exit!")
}
