package main

import (
	"fmt"
)

type Render interface{
	Render()
}

type Button struct {}

func (b Button) Render(){
	fmt.Println("Button")
}
type Window struct{}

func (b Window) Render(){
	fmt.Println("Window")
}
func render(r Render){
	r.Render()
}

/*func main(){
	var r Render
	r = Button{} //етот варіант похож на действия компелятора во время render(Button{})
	render(r)
	r = Window{}
	render(r)
	//render(Button{})
	//render(Window{})
}

//Интерфуйси - набор методов, которие надо реализовать, чтоби удовлетварить интерфейсу.
//Ключевое слово interface
// type Stringer interface{
//	String() string }

//Переменная типа интерфейс может сожержать 
//значение типа, реализуещего етот интерфейс.
// var s Stringer
// s = time.Time{}*/