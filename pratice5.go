package main

import "fmt"

func main() {
	s1 := []int{2,4,6,8}
	s2 := make([]int,16)
	//s3 := make([]int,16,32)

	fmt.Println("Copying slice ")
	copy(s2,s1)
	fmt.Println(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3],s2[4:]...)
	fmt.Printf("s2 %v,len %d,cap %d\n",s2,len(s2),cap(s2))

	fmt.Println("Poping from front")
	s2 = s2[1:]
	fmt.Printf("s2 %v,len %d,cap %d\n",s2,len(s2),cap(s2))

	fmt.Println("Poping from tail")
	s2 = s2[:len(s2)-1]
	fmt.Printf("s2 %v,len %d,cap %d\n",s2,len(s2),cap(s2))

}
