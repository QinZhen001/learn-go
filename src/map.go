package main

import "fmt"

func create_map() {
	fmt.Println("============ create_map ===============")
	// key in map 无序的
	m := map[string]string{
		"name": "rick",
		"site": "xxx",
	}

	m2 := make(map[string]int) // empty map
	var m3 map[string]int      // nil

	fmt.Println(m, m2, m3)

	fmt.Println("======= traverse map ==============")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("============= get value ================")
	name, ok := m["name"]
	fmt.Println("get name value: ", name, ok)

	name1, ok := m["xxxxx"]
	fmt.Println("get does not exist value: ", name1, ok) // zero value  空

	if _, ok := m["xxxxx"]; ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}

	fmt.Println("========== delete value =======================")
	delete(m, "name")
	fmt.Println(m)

}
