package main

import "fmt"

type usb interface {
	Name() string
	Connect
}

type Connect interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect "+pc.name)
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

type TvConnector struct {
	name string
}

func (tv TvConnector) Connect() {
	fmt.Println("Connect "+tv.name)
}

func Disconnect(u interface{})  {
	switch  v:=u.(type){
	case PhoneConnector:
		fmt.Println("Disconnect "+v.name)
	case TvConnector:
		fmt.Println("Disconnect "+v.name)
	default:
		fmt.Println("Unknown device")
	}

}

func main() {
	//var p usb
	//p = PhoneConnector{"phoneconnector"}
	//p.Connect()
	//Disconnect(p)

	//var c Connect
	//c = PhoneConnector{"phoneconnector"}
	//c.Connect()
	//Disconnect(c)

	tv := TvConnector{"TvConnector"}
	tv.Connect()
	Disconnect(tv)
}
