package main

import (
	"fmt"
)

type BSTree struct {
	root *Node
}

type Node struct {
	left *Node
	right *Node
	value int
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

//查看树中是否具有某个元素,查
func (bst *BSTree) SearchBST(value int) bool {
	node := bst.root
	for node != nil{
		if node.value == value{
			return true
		}else if value < node.value{
			node = node.left
		}else {
			node = node.right
		}
	}
	return false
}

//向树中增加元素
func (bst *BSTree) Insert(value int) {

	if bst.root == nil{
		bst.root = &Node{value:value}
		return
	}
	if bst.SearchBST(value) == true{
		fmt.Println("该数据已存在于树中")
		return
	}
	x := bst.root
	var parent *Node
	for x!=nil{
		parent = x
		if value < x.value{
			x = x.left
		}else {
			x = x.right
		}
	}
	if value < parent.value{
		parent.left = &Node{value:value}
	}else {
		parent.right = &Node{value:value}
	}
}

//PrintTree1 递归结构，中序遍历树
//func (t *Node) PrintTree1() {
//	if t.left != nil {
//		t.left.PrintTree1()
//	}
//	fmt.Print(t.value," ")
//	if t.right != nil {
//		t.right.PrintTree1()
//	}
//}

func (t *Node) PreOderTraverse() {
	fmt.Print(t.value," ")
	if t.left != nil {
		t.left.PreOderTraverse()
	}
	if t.right != nil {
		t.right.PreOderTraverse()
	}
}

func (t *Node) PostOderTraverse() {

	if t.left != nil {
		t.left.PostOderTraverse()
	}
	if t.right != nil {
		t.right.PostOderTraverse()
	}
	fmt.Print(t.value," ")
}

func (t *Node) InOderTraverse() {

	if t.left != nil {
		t.left.InOderTraverse()
	}
	fmt.Print(t.value," ")
	if t.right != nil {
		t.right.InOderTraverse()
	}

}

//寻找树中的最小值
func (bst *BSTree) Min() int  {
	node := bst.root
	for node.left != nil{
		node = node.left
	}
	return node.value
}

func (bst *BSTree) Max() int  {
	node := bst.root
	for node.right != nil{
		node = node.right
	}
	return node.value
}

func (bst *BSTree) IsEmpty() bool {
	if bst.root == nil{
		return true
	}
	return false
}

//通过树节点的值去修改它的值
func (bst *BSTree)Change(oldvalue int,newvalue int) bool {
	node := bst.root
	for node != nil{
		if node.value == oldvalue{
			node.value = newvalue
			return true
		}else if oldvalue < node.value{
			node = node.left
		}else {
			node = node.right
		}
	}
	return false
}

func main() {
	var bst *BSTree
	bst = NewBSTree()	//初始化
	fmt.Println(bst.root)
	//bst.Insert(12)
	//bst.Insert(5)
	//bst.Insert(3)
	//bst.Insert(8)
	//bst.Insert(20)
	//bst.Insert(18)
	//bst.Insert(30)
	bst.Insert(1)	//插入元素
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(0)
	bst.Insert(10)
	bst.Insert(11)
	fmt.Println(bst.root)
	fmt.Println(bst.SearchBST(10))	//查找值为10的树是否存在于二叉树中
	bst.root.PreOderTraverse()				//前序遍历
	fmt.Println()
	//fmt.Println()
	//fmt.Println(bst.SearchBST(0))
	//fmt.Println(bst.Min())
	//fmt.Println(bst.Max())
	//fmt.Println(bst.root.value)
	//fmt.Println(bst.root.left.value)
	//fmt.Println(bst.root.right.value)
	bst.root.InOderTraverse()				//中序遍历
	fmt.Println()
	bst.root.PostOderTraverse()				//后序遍历
	bst.Change(10,20)
	fmt.Println()
	bst.root.InOderTraverse()


}
