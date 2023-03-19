package queue

import "fmt"

func ExampleQueue_Push() {
	q := Queue{1, 2, 3}
	q.Push(4)
	q.Push(5)
	for !q.IsEmpty() {
		fmt.Println(q.Pop())
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}
