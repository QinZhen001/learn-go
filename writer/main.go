package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			// EOF, err == io.EOF
			fmt.Println(err)
			break
		}
		println(string(p[:n]))
	}

	meats := []string{
		"pork",
		"beef",
		"chicken",
	}
	var writer bytes.Buffer
	for _, meat := range meats {
		n, err := writer.Write([]byte(meat))
		if err != nil {
			fmt.Println(err)
			break
		}
		if n != len(meat) {
			fmt.Println("failed to write data")
			break
		}
		fmt.Println("write data:", writer.String())
	}
}
