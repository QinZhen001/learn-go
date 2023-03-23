package main

import (
	"fmt"
	"time"
)

// chan 是一个引用类型，它的零值是 nil ，它必须使用 make 来初始化。
// chan 是goroutine之间的通信机制，它是一个数据结构，是一个队列，遵循先入先出的原则。
// chan 发数据，必须有接收者，否则会阻塞。

func worker(c <-chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func worker1(id int, c chan int) {
	// for {
	// 	// close 这里还会收到数据 初始值
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("Worker %d received %d\n",
	// 		id, n)
	// }

	for n := range c {
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}

}

func chanDemo() {
	c := make(chan int)
	go worker(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go worker(c)
	return c
}

func chanDemo1() {
	var channels [10]chan<- int // 只能发
	for i := 0; i < len(channels); i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < len(channels); i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) // 缓冲区
	go worker1(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker1(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	// chanDemo1()
	// bufferedChannel()
	channelClose()
}
