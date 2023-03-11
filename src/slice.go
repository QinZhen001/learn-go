package main

import "fmt"

// [xxx:xxx] 是一个slice  （slice是对array的一个视图 view）  （slice本身是没有数据的）
// slice底层结构： ptr指向开始 len长度 cap（包括len）  =>  只要不超过cap slice都可以向后扩展，但不能向前扩展
func slice() {
	fmt.Println("========= slice =================")
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] => ", arr[2:6])
	fmt.Println("arr[:6] => ", arr[:6])
	fmt.Println("arr[2:] => ", arr[2:])
	fmt.Println("arr[:] => ", arr[:])

	s1 := arr[2:6]
	updateSlices(s1)
	fmt.Println(s1)
	fmt.Println(arr) // arr 也会被改变

	fmt.Println("-------- Reslice --------------")
	s2 := s1[2:4]
	fmt.Println(s2)
	s2 = s2[1:2]
	fmt.Println(s2)

	fmt.Println("----------  Extending slice")
	s3 := arr[2:6]
	s4 := s3[3:5] // 可以取到
	// fmt.Println(s3, s4)
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v,len(s4)=%d,cap(s4)=%d\n", s4, len(s4), cap(s4))

	fmt.Println("========= append ==============")
	s5 := append(s4, 10) // 可以改到原数组arr元素
	s6 := append(s5, 11) // 不可以改到原数组arr元素 =>  已超过arr最后坐标 > view different arr
	fmt.Println("s5,s6= ", s5, s6)
	fmt.Println("arr= ", arr)
}

// int[] 就是个slice
func updateSlices(s []int) {
	fmt.Println("========== updateSlices ========")
	s[0] = 1000
}

func create_slice() {
	fmt.Println("================== create_slice ================")
	// zero value for slice is nil
	var s []int
	for i := 0; i < 20; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	fmt.Println("====================")
	s1 := []int{2, 4, 6, 8} // create slice
	s2 := make([]int, 16)   // know length
	s3 := make([]int, 16, 32)
	println(s1, s2, s3)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("=========== copy slice ================")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("================ delete element from slice ====================")
	// 删除3 index 元素
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("============== popping from front ================")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("============== popping from back ================")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}
