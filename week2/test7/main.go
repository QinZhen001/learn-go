package main

import "fmt"

func GivenFood() chan string {
	ch := make(chan string)
	go func() {
		ch <- "pizza"
		ch <- "hamburger"
		ch <- "hotdog"
		close(ch)
	}()

	return ch
}

func main() {
	ch := GivenFood()

	for {
		if input, ok := <-ch; ok {
			println(input)
		} else {
			break
		}
	}

	fmt.Println("end")

	// for data := range ch {
	// 	fmt.Println(data)
	// }
}
