package main

import (
	"fmt"
	"math"
)

func eval(a,b int,op string ) (int,error) {
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		return a/b,nil
	default:
		return 0,fmt.Errorf("unsupprted operator %s",op)
	}
}
func div(a,b int) (q,r int) {
	return a/b,a%b
}

func apply(op func(int,int) int,a,b int) int {

	return op(a,b)
}

func sum(numbers ...int)  int{
	s :=0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}

func main() {
	q,r := div(13,3)
	fmt.Println(q,r)
	fmt.Println(eval(13,3,"+"))
	if result, err := eval(12,4,"#") ; err != nil{
		fmt.Println("Error:",err)
	}else{
		fmt.Println(result)
	}

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a),float64(b)))
	},3,4))

	fmt.Println(sum(1,2,3,4))
}
