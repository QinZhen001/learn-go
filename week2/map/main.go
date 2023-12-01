package main

func main() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	value, exists := myMap["a"]
	if exists {
		println(value)
	}
}
