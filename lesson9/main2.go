package main

import(
	"fmt"
)

type Duck interface{
	Talk() string
	Walk()
	Swim()
}

type Dog struct{
	name string
}

func (d Dog) Talk() string{
	return "AGGGRRR"
	//собака умеет говорить
}

func (d Dog) Walk(){
//собака умеет гулять
}

func (d Dog) Swim(){
//собака теперь умеет плавать
}
func quack(d Duck) {
	fmt.Println(d.Talk())
}
//можна считать что собака теперь утка, поскольку удолетваряет интерфейс
/*func main(){
	quack(Dog{})
}*/
//Не надо писать как в Java или в PHP имплемент что то или какой то интерфейс
//ми реализуем какието методи и go понимает что есть реализация етих методов и что стуктура удолетваряет определеному интерфейсу
