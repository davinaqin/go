package main

import "fmt"

type ArrayList struct {
	size int
	Elements []interface{}
}

//初始化
func New(value ...interface{}) *ArrayList {
	list := &ArrayList{}
	list.Elements = make([]interface{},10)
	if len(value) >0{
		list.Add(value...)
	}
	return list
}

//增加
func (list *ArrayList) Add(value ...interface{})  {
	if list.size+len(value) > len(list.Elements){//为什么要大于等于开辟空间-1，而不是大于等于开辟空间
		newElements := make([]interface{},list.size+len(value))
		copy(newElements,list.Elements)
		list.Elements = newElements
	}

	for _,value := range value{
		list.Elements[list.size] = value
		list.size++
	}
}

//删
func (list *ArrayList) Remove(index int) interface{} {
	if index <0 || index >= list.size{
		return nil
	}
	curEle := list.Elements[index]
	list.Elements[index] = nil
	copy(list.Elements[index:],list.Elements[index+1:list.size])
	list.size--
	return curEle
}

//查
func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= list.size{
		return nil
	}
	return list.Elements[index]
}

//改
func (list *ArrayList) Change(index int,value interface{})  {
	if index <0 || index >= list.size{
		return
	}
	list.Elements[index] = value
}

func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayList) Size() int  {
	return list.size
}

func (list *ArrayList) Contains(value interface{}) bool {
	for _,curValue := range list.Elements{
		if value == curValue{
			return true
		}
	}
	return false
}

//遍历
func (list *ArrayList) Traverse()  {
	for i:=0 ;i<list.size;i++{
		fmt.Printf("%v " ,list.Elements[i])
	}
	fmt.Println()
}

func main()  {
	arrlist := New(1,2,3)
	arrlist.Traverse()
	arrlist.Add(4,5,6,7,8,9,10,11)
	fmt.Println(arrlist)
	arrlist.Remove(6)
	fmt.Println(arrlist)
	fmt.Println(arrlist.Get(2))
	fmt.Println(arrlist.IsEmpty())
	fmt.Println(arrlist.Size())
	fmt.Println(arrlist.Contains(15))
	arrlist.Traverse()
}
