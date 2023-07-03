package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 8)

	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}

	fmt.Println("=====================================")

	meats := []string{
		"calf", "pig", "chicken", "duck", "cow",
	}
	var writer bytes.Buffer
	for _, meat := range meats {
		n, err := writer.Write([]byte(meat))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(meat) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
		fmt.Println(writer.String())
	}
	fmt.Println(writer.String())
}
