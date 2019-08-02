package main

type Inter interface {
	Ping()
	Pong()
}

type St struct {

}

func (St) Ping() {
	println("ping")
}

func (*St) Pong() {
	println("pong")
}

func main() {
	//var st = St{}
	//st.Pong()
	//st.Ping()
	//var p Inter = &st
	//p.Pong()

	var st *St = nil
	var p Inter = st
	p.Pong()
}
