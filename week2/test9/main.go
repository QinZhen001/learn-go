package main

import "fmt"

func getchan(i int, chs []chan string) chan string {
	fmt.Printf("chs: %d", i)
	return chs[i]
}

func main() {
	ch4 := make(chan string, 8)
	ch5 := make(chan string, 8)
	var chs = []chan string{ch4, ch5}
	select {
	case getchan(0, chs) <- "name1":
		fmt.Println("执行了第一条语句")
	case getchan(1, chs) <- "name2":
		fmt.Println("执行了第二条语句")
	default:
		fmt.Println("默认")
	}

	// ch:= make(chan string)
	// go func ()  {

	// }
}
