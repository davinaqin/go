package main

import (
	"fmt"
	"math"
)

func consts()  {
	const  a, b  =3, 4
	const filename  = "abc.txt"
	var c int
	c = int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c)
}

func enums()  {
	const(
		cpp = iota
		java
		python
		golang
		javascript
		)
	const(
		b = 1 <<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang,javascript)
	fmt.Println(b,kb,mb,gb,tb,pb)
}
func main() {
	consts()
	enums()
}
