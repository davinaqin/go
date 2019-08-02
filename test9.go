package main

import "fmt"

func updateSlice(s []int)  {
	s[0] = 100
}

func change(arr []int)  {
	arr[0] = 200
	for i,v := range arr{
		fmt.Println(i,v)
	}
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6}
	change(arr[:])
	s1 := arr[2:6]
	s2 := arr[:]
	fmt.Println("arr[2:6]:",arr[2:6])
	fmt.Println("arr[:]:",arr[:])

	updateSlice(s1)
	fmt.Println("after update s1",s1)
	updateSlice(s2)
	fmt.Println("after update s2",s2)

	fmt.Println("arr",arr)

	fmt.Println("Extending slice")
	s3 := arr[2:6]
	s4 := s1[3:5]
	fmt.Printf("s3 is %v ,len(s3) is %d,cap(s3) is %d\n",s3,len(s3),cap(s3))
	fmt.Printf("s4 is %v ,len(s4) is %d,cap(s4) is %d",s4,len(s4),cap(s4))

}
