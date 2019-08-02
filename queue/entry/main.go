package main

import (
	"awesomeProject/queue"
	"fmt"
)
func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
