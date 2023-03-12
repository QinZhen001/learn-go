package main

import (
	"fmt"
	"unicode/utf8"
)

func create_string() {
	fmt.Println("========= create_string =============")
	s := "yes阿萨德" // 每个中文三字节 UTF-8
	fmt.Printf("%s len=%d \n", s, len(s))
	fmt.Printf("%X\n", s)

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Printf("\n")

	for i, ch := range s {
		// ch is rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Printf("\n")

	fmt.Println("rune count: ", utf8.RuneCountInString((s)))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Printf("\n")

	// 如果想知道第几个字符到底是什么 => rune
	// decode 出来的结果 每个字符4字节
	// []rune 再开一个内存空间
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Printf("\n")
}
