package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	fmt.Println("========= convertToBin ==========")
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile1(filename string) {
	fmt.Println("========= printFile1 =============")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// 省略初始条件 和 递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
