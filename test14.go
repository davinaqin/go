package main

import "fmt"

type TreeNode struct {
	value int
	Left,Right *TreeNode
}

//使用了自定义工厂
//注意返回了局部变量地址
func createNode(value int) *TreeNode {
	return &TreeNode{value:value}
}

//(node TreeNode )表示函数的接收者
func (node TreeNode) print()  {
	fmt.Println(node.value)
}

func (node *TreeNode) setValue(value int)  {
	node.value = value
}

func (node *TreeNode) traverse()  {
	if node == nil{
		return
	}
	node.Left.traverse()
	node.print()
	node.Right.traverse()
}

func main() {
	var root TreeNode
	root = TreeNode{value:3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5,nil,nil}
	root.Right.Left = new(TreeNode)
	root.Right.Left.setValue(4)

	//nodes := [3]TreeNode{
	//	{},
	//	{},
	//	{},
	//}
	//fmt.Println(nodes[0],nodes[1],nodes[2])

	//root.setValue(20)
	//root.print()

	root.traverse()

}
