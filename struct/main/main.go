package main

// 一个目录只能一个package
// 所以package main 要在创一个目录

import (
	"fmt"
	"learngo/struct/tree"
)

func main() {
	fmt.Println("============= create_struct ===========")
	var root tree.TreeNode
	fmt.Println(root)

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.TreeNode)
	root.Right.Right = tree.CreateNode(100)
	fmt.Println(root)

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{Value: 6, Left: nil, Right: nil},
	}
	fmt.Println(nodes)

	root.Right.Left.SetValue(99)
	root.Right.Left.Print()

	fmt.Println()

	root.Traverse()

	fmt.Println()
}
