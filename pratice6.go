package main

import "fmt"

func main() {
	m := map[string]string{
		"name":"qinlifei",
		"course":"golang",
		"site":"immoc",
		"quality":"notbad",
	}
	m1 := make(map[string]int)	//m2 ==  empty map
	var m2 map[string]int	//m3 == nil
	fmt.Println("m",m)
	fmt.Println("m1",m1)
	fmt.Println("m2",m2)

	//遍历
	for k,v := range m{
		fmt.Printf("key %v,value %v\n",k,v)
	}

	//取值
	for _,v := range m{
		fmt.Println(v)
	}
	courseName := m["course"]
	fmt.Println(courseName)

	if courseName,ok := m["qin"];ok{
		fmt.Println(courseName)
	}else{
		fmt.Println("Key does not exit")
	}

	delete(m,"name")
	fmt.Println("After delete name",m)

	fmt.Println("Map length",len(m))

}
