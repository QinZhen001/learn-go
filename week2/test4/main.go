package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteMenu(fileName string, s []string) {
	// 当前执行文件的路径
	curDir, _ := os.Getwd()
	path := curDir + fileName
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, item := range s {
		fmt.Fprintln(w, item)
	}
}

func main() {

	s := []string{
		"calf", "pig", "chicken", "duck", "cow",
	}
	WriteMenu("/menu.txt", s)
}
