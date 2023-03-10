package main

import "fmt"

// 指针
// 指针不能运算
// go 语言只有值传递一种方式 （拷贝）（如果是指针，拷贝地址）
// 定义数据的时候就要考虑当做一个值来使用 还是 当做一个指针来使用

// 错误的写法 （值传递，无法交换）
// func swap(a, b int) {
// 	b, a = a, b
// }

func swap(a, b *int) {
	fmt.Println("=========== swap ================")
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	fmt.Println("================= swap1 ==================")
	return b, a
}
