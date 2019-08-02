package main

import "fmt"

func main() {
	//可以在声明变量并赋值的语句中，省略变量的类型部分，go语言可以推导出该变量的类型
	var num2 = 5.89E-4
	fmt.Printf("浮点数 %E 表示的是 %f。\n",num2,float64(num2))
	var num3 = 3.7E+1 + 5.9E-2i
	fmt.Printf("浮点数 %E 表示的是 %f。\n",num3,complex64(num3))
	var char rune = '赞';
	fmt.Printf("字符 '%c' 的Unico代码点是 %s。\n",char,rune(char))
}
