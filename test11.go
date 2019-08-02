package main

import "fmt"

func main() {
	m := map[string]string {
		"name":"ccmouse",
		"course":"golang",
		"site":"immoc",
		"quality":"notbad",
	}
	m2 := make(map [string]int)

	var m3 map[string]int

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k,v := range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courseName,ok := m["course"]
	fmt.Println(courseName,ok )
	//causeName,ok := m["cause"]
	//fmt.Println(causeName,ok)

	if causeName,ok := m["cause"];ok{
		fmt.Println(causeName,ok)
	}else {
		fmt.Println(causeName,ok)
	}

	fmt.Println("delete values")
	name,ok := m["name"]
	fmt.Println(name,ok)

	delete(m,"name")
	name,ok = m["name"]
	fmt.Println(name,ok)

}
