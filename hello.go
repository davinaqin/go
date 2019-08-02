package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence)*damageRate
	fmt.Println(damage)
	hp := 100
	fmt.Println(hp)

	var a int = 100
	var b int = 200
	a,b = b,a
	fmt.Println(a,b)
}
