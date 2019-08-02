package main

import (
	"fmt"
	"math"
)

type Point struct {
	x,y int
}

type Polar struct {
	x,y,z int
}

func Abs(p Point) int {
	return int(math.Hypot(float64(p.x),float64(p.y)))
}

func main() {
	p1 := Point{3,4}
	i := Abs(p1)
	fmt.Println(i)
}
