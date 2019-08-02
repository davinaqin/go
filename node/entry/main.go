package main

import (
	"awesomeProject/node"
	"fmt"
)

type MyTreeNode struct {
	node *node.TreeNode
}

func (myNode *MyTreeNode) posOrder()  {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.posOrder()
	right := MyTreeNode{myNode.node.Right}
	right.posOrder()
	myNode.node.Print()
}

func main() {
	var root node.TreeNode
	root = node.TreeNode{Value:3}
	root.Left = &node.TreeNode{}
	root.Right = &node.TreeNode{5,nil,nil}
	root.Right.Left = new(node.TreeNode)
	root.Right.Left.SetValue(4)

	//nodes := [3]TreeNode{
	//	{},
	//	{},
	//	{},
	//}
	//fmt.Println(nodes[0],nodes[1],nodes[2])

	//root.setValue(20)
	//root.print()

	root.Traverse()

	fmt.Println()
	myRoot := MyTreeNode{&root}
	myRoot.posOrder()
}
