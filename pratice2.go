package main

import "fmt"

func main() {
	var numbers = [...]int{1,2,3,4,5,6,7,8,9,10}
	slice5 := numbers[4:6:8] //[5,6]，不能访问第九位元素
	length := len(slice5)
	capacity := cap(slice5)
	fmt.Printf("slice %v ,slice5 len %d,slice5 cap %d\n",slice5,length,capacity)
	slice5 = slice5[:cap(slice5)]	//把slice5扩展到cap的长度
	slice5 = append(slice5,11,12,13)
	fmt.Printf("slice5 %v len %d\n",slice5,len(slice5))
	slice6 := []int{0,0,0}
	copy(slice5,slice6)
	fmt.Printf("%v,%v,%v\n",slice5[2],slice5[3],slice5[4])

}
