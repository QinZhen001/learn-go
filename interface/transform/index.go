package transform

import "fmt"

func Interface2String(inter interface{}) {
	switch inter.(type) {
	case string:
		fmt.Println("string: ", inter.(string))
	case int:
		fmt.Println("int: ", inter.(int))
	case float64:
		fmt.Println("float64: ", inter.(float64))
	default:
		println("unknown")
	}
}
