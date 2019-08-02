package main

import "fmt"

//死循环
func forever()  {
	for{
		fmt.Println("abc")
	}
}

func main() {
	sum := 0
	for i :=1 ;i <=100 ;i++{
		sum += i
	}
	fmt.Println(sum)
}
