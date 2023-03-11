package main

import "fmt"

func test_arr1() {
	fmt.Println("============== test_arr1 ===================")
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 4, 6, 8, 0} // ... 代替定义具体数字

	var grid [4][5]int

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// [3]int 和 [5]int 不是一个类型
	// printArr(arr2)

	fmt.Println("================")
	printArr(arr1)
	fmt.Println(arr1)
	fmt.Println("====================")
	printArr1(&arr1)
	fmt.Println(arr1)
	fmt.Println("====================")
	printArr2(arr2[:]) // 传进去 slice
	fmt.Println(arr2)
}

// 数组是值类型 （会拷贝数组） =>  go 语言一般不直接使用数组 使用slice切片
func printArr(arr [5]int) {
	arr[0] = 100 // 影响不到原类型
	fmt.Println("============ printArr ======================")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArr1(arr *[5]int) {
	arr[0] = 100 // 影响不到原类型
	fmt.Println("============ printArr1 ======================")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArr2(arr []int) {
	arr[0] = 100 // 影响不到原类型
	fmt.Println("============ printArr2 slice ======================")
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
