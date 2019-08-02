package main

import "fmt"

type T struct {
	i1 int
	f1 float32
	str string
}

type Interval struct {
	start int
	end int
}

func main() {
	ms := new(T)
	ms.i1 = 0
	ms.f1 = 0.24
	ms.str = "string"
	fmt.Println(ms.i1,ms.f1,ms.str)
	fmt.Println(ms)

	var ms1 *T
	ms1 = &T{1,0.3,"myname"}
	fmt.Println(ms1)

	ms2 := &T{2,0.1,"mybook"}
	fmt.Println(ms2)

	var ms3 T
	ms3 = T{3,0.9,"mycomputer"}
	fmt.Println(ms3)

	intr1 := Interval{1,2}
	intr2 := Interval{end:3,start:1}
	intr3 := Interval{end:4}

	fmt.Println(intr1,intr2,intr3)
}
