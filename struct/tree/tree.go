package tree

import (
	"fmt"
)

// go 仅支持封装 不支持继承和多态
// go 没有 this 指针

// 要改变内容必须采用指针接收者
// 结构体过大也考虑指针接收者

// 针对包 （每个目录只能有一个包）
// 首字母大写 public
// 首字母小写 private

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 这里的 node 也是值传递 (拷贝一份出来)
func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

// func (node TreeNode) SetValue(v int) {
// 	// 因为node是传值 所以这里改不掉
// 	node.Value = v
// }

func (node *TreeNode) SetValue(v int) {
	if node == nil {
		fmt.Println("set node Value nil")
		return
	}
	node.Value = v
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// go 没有构造函数
// 可以自定义工厂函数
func CreateNode(value int) *TreeNode {
	// 局部变量也可以返回给外部用
	// go 不需要知道结构创建在堆还是栈上
	return &TreeNode{Value: value}
}
