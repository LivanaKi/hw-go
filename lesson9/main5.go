package main

import(
	"fmt"
)

type Speaker interface{
	SayHello()
	String()
}

type Human struct{
	Greeting string
}

func (h *Human) SayHello(){
}
func (h Human) String(){fmt.Println(h.Greeting)}
/*
func main() {
	h := Human{"Hello"}
	var s Speaker = &h
	s.String()
}

/*func meet(s Speaker){
	s.SayHello()
} 

func main(){
	var s Speaker = Human{Greeting: "Hello"}
	//можна посмотреть динамический тип и значение
	fmt.Printf("%v | %T", s, s)
}*/