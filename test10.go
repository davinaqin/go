package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("%v,len=%d,cap=%d\n",s,len(s),cap(s))
}

func main() {
	//arr := [...]int{0,1,2,3,4,5,6,7}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//s3 := append(s2,10)
	//s4 := append(s3,11)
	//s5 := append(s4,12)


	var s []int //Zero value for slice is nil
	for i:= 0;i<100;i++{
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{1,2,3,4}
	s2 := make([]int,16)
	s3 := make([]int,10,32)

	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copying Slice")
	copy(s2,s1)
	printSlice(s2)

	fmt.Println("deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Poping from front")
	front := s2[1:]
	s2 = s2[1:]
	printSlice(front)

	fmt.Println("Poping from back")
	tail := s2[:len(s2)-1]
	s2 = s2[:len(s2)-1]
	printSlice(tail)
}
