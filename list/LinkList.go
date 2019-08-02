package main

import "fmt"

type LinkList struct {
	Length int
	Head *Node
	Tail *Node
}
type  Node struct{
	Data int
	Next *Node
}

//初始化链表
func (list *LinkList) Init()  {
	list.Length = 0
	list.Head = nil
	list.Tail = nil
}

//在链表末尾加入元素
func (list *LinkList) Append(node *Node) {
	if list.Length == 0{
		list.Head = node
		list.Tail = node
	}else {
		list.Tail.Next = node
		list.Tail = node
	}
	list.Length ++
}

//通过索引获取Node节点，查
func (list *LinkList) Get(i int) *Node {
	if list.Length == 0 || i >= list.Length || i<0{
		fmt.Println("This node is not exist")
		return nil
	}else {
		p := list.Head
		for j:=0;j<i;j++{
			p = p.Next
		}
		return p
	}
}

//遍历链表
func (list *LinkList) Traverse() {
	point := list.Head
	for i := 0;i<list.Length;i++ {
		fmt.Printf("%d ",point.Data)
		point = point.Next
	}
	fmt.Println()
}

//删除链表中的第i个节点，删
func (list *LinkList) Delete(i int)  {
	if list.Length <= 0 || i > list.Length{
		fmt.Println("Empty LinkList")
	}else {
		p := list.Head
		for j:=0 ; j< i-2 ;j++{
			p = p.Next
		}
		p.Next = p.Next.Next
		fmt.Println("Delete node ",i)
		list.Length--
	}
}

//在指定位置前插入元素，增
func (list *LinkList) Insert(index int,node *Node) bool{
	if list.Length == 0{
		fmt.Println("Empty LinkList")
		return false
	}else if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Length++
	}else if index == list.Length {//在最后一个元素的后边插入元素
		list.Tail.Next = node
		list.Tail = node
		list.Length++
	}else {
		pre := list.Head
		for j:=1 ; j< index;j++{
			pre = pre.Next
		}
		node.Next = pre.Next
		pre.Next = node
		list.Length++
	}
	return true
}

//改
func (list *LinkList) Change(pos int,value int) {
	point := list.Head
	for i:=0 ;i< pos-1;i++{
		point = point.Next
	}
	point.Data = value
}



func main()  {
	l := LinkList{}
	l.Init()
	for i := 1 ;i<=10;i++{
		l.Append(&Node{i,nil})
	}
	fmt.Printf("Traverse LinkList:")
	l.Traverse()
	l.Insert(10,&Node{45,nil})
	l.Traverse()
	l.Delete(6)
	l.Traverse()
	fmt.Println(l.Get(6).Data)
	l.Change(9,20)
	l.Traverse()
}
