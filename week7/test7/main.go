package main

// only write to channel
func product(ch chan<- int) {
	for {
		ch <- 1
	}
}

// only read from channel
func consume(ch <-chan int) {
	for {
		<-ch
	}
}

func main() {
	var c = make(chan int)
	go product(c)
	go consume(c)
}
