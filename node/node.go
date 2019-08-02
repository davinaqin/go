package node

import "fmt"


type TreeNode struct {
	Value int
	Left,Right *TreeNode
}

//使用了自定义工厂
//注意返回了局部变量地址
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value:value}
}

//(node TreeNode )表示函数的接收者
func (node TreeNode) Print()  {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int)  {
	node.Value = value
}

func (node *TreeNode) Traverse()  {
	if node == nil{
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
