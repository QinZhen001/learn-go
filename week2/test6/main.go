package main

import "runtime"

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// 把代码中的runtime.Gosched()注释掉，执行结果是：
// hello
// hello
// 因为say("hello")这句占用了时间，等它执行完，线程也结束了，say("world")就没有机会了。

// 没有在代码中通过
// runtime.GOMAXPROCS(n) 其中n是整数，
// 指定使用多核的话，goroutins都是在一个线程里的，它们之间通过不停的让出时间片轮流运行，达到类似同时运行的效果。
