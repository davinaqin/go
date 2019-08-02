package main

import "fmt"

func swap(a,b *int) (int,int) {
	 *a , *b = *b ,*a
	 return *a,*b
}


func main() {
	var a int = 2
	var b int = 4
	var pa *int = &a
	//*pa += 1
	fmt.Println(*pa)
	a,b = swap(&a,&b)
	fmt.Println(a,b)

}
