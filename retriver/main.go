package main

import (
	"awesomeProject/retriver/mock"
	"fmt"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriver
	r = mock.MockRetriver{"this is a fake imocc.com"}
	s := download(r)
	fmt.Println(s)
}
