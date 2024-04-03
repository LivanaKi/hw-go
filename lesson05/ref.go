package main

import(
	"fmt"
	"reflect"
)
/*
type User struct{
	Name string
	age int
}

func main(){
	u := User{"Vasya", 50}
	age := reflect.ValueOf(&u).Elem().FieldByName("age")
	fmt.Println(age.Int())
}