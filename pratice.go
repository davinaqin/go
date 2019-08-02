package main

import "fmt"

func main() {
	arra := [...]int{3,4,1,5}
	maxi := -1
	maxValue := -1
	//练习求数组的最大值
	for i,v := range arra{
		if v > maxValue{
			maxi = i
			maxValue = v
		}
	}
	fmt.Println(maxi,maxValue)



}
