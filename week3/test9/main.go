package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// exit := make(chan os.Signal)
	// signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	// <-exit
	// fmt.Println("程序优雅退出结束")

	sigs := make(chan os.Signal, 1)
	// signal.Notify registers the given channel to receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	// 不影响后面代码执行
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
