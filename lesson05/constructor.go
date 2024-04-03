package main

import(
	"fmt"
)

type countriseCodes map[string]int

type world struct{
	Name string
	countrise countriseCodes
}

func NewWorld(name string) *world{
	return &world{
		Name: name,
		countrise: make(countriseCodes),
	}
}

func (w *world) AddCountry(name string, code int){
	w.countrise[name] = code
}

func (w world) GetContryCode(name string) (code int, ok bool){
	code, ok = w.countrise[name]
	return
}
/*
func main(){
	w := NewWorld("Earth")
	w.AddCountry("Ukraine", 1142)
	c, _ := w.GetContryCode("Ukraine")
	fmt.Println(c)
}