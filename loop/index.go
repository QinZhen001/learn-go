package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	s := `
	aaa
	bbb
	ccc
	ddd`
	printFileContents(strings.NewReader(s))
}
