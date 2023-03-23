package main

import (
	"fmt"
	"time"
)

// goroutine 其实是协程，是一种比线程更加轻量级的存在，它的调度是由 Golang 运行时进行管理的。
// 非抢占式多任务处理，由协程主动交出控制权。
// 编译器/解释器/虚拟机层面的多任务。
// 不是操作系统层面的多任务。
// 调度器是运行时的一部分，它是一个单线程的循环，不断地在所有的可运行的 goroutine 中做任务调度。
// goroutine 切换点 1. I/O，2. channel，3. 等待锁，4. 函数调用（有时）
// channel 本身就是 goroutine 之间的通信机制，它是一个数据结构，是一个队列，遵循先入先出的原则。

// main 函数也是个 goroutine
func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//
	// var a [10]int
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		for {
	// 			a[i]++
	// 			// 手动交出控制权 才能停下来
	// 			runtime.Gosched()
	// 		}
	// 	}(i)
	// }
	// // 停不下来 被一个携程霸占了
	// time.Sleep(time.Millisecond)
	// fmt.Println(a)
}
