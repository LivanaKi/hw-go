package main

import "fmt"

type Hound interface{
	Hunt()
}

type Poodle interface {
	Bark() string
}

type Decorative interface{
	Sleep(hours int)
}

type Bloodhound struct { name string}
func (b Bloodhound) Hunt() { fmt.Printf("%s is hunting\n", b.name)}

type TeacupPoodle struct { name string}
func (p TeacupPoodle) Bark() string{ return fmt.Sprintf("%s is barking...\n", p.name)}

type ShihTzu struct{ name string}
func (p ShihTzu) Sleep(horts int){fmt.Printf("%s is sleeping...\n", p.name)}

func zoo(dog interface{}){
	switch d := dog.(type){
	case Hound:
		d.Hunt()
	case Poodle:
		fmt.Print(d.Bark())
	case Decorative:
		d.Sleep(3)
	default:
		fmt.Println("Unknown")
	}
}
func main(){
	zoo(Bloodhound{"misha"})
	zoo(TeacupPoodle{"vasya"})
	zoo(ShihTzu{"Jacky"})
}