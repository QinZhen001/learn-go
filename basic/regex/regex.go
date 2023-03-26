package main

import (
	"fmt"
	"regexp"
)

const text = `my email is xxx@gmail.com
email1 is bbb@gmail.com
email2 is ccc@gmail.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.[a-zA-Z0-9]+`)
	// match := re.FindAllString(text, -1)
	// fmt.Println(match)

	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
