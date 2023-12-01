package main

import "fmt"

type ParameterStruct struct {
	Name string
}

func changeParameter(para *ParameterStruct, value string) {
	para.Name = value
}

func cannotChangeParameter(para ParameterStruct, value string) {
	para.Name = value
}

func main() {
	str := "hello world"
	pointer := &str
	anotherString := *&str
	fmt.Println(str)
	fmt.Println(pointer)
	fmt.Println(anotherString)
	fmt.Println("--------------------------")
	str = "changed string"
	fmt.Println(str)
	fmt.Println(pointer)
	fmt.Println(anotherString)
	fmt.Println("--------------------------")
	para := ParameterStruct{Name: "aaa"}
	fmt.Println(para)
	changeParameter(&para, "bbb")
	fmt.Println(para)
	cannotChangeParameter(para, "ccc")
	fmt.Println(para)
}
