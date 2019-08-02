package main

import "fmt"

type Rectangle struct {
	x int
	y int
}

func Area(r Rectangle) int {
	return r.x*r.y
}

func Perimeter(r Rectangle) int {
	return (r.x+r.y)*2
}

func main() {
	r1 := Rectangle{3,4}
	area := Area(r1)
	pre := Perimeter(r1)
	fmt.Println(area,pre)
}
