package main

import(
	"fmt"
)
/*
type User struct{
	Id int64
	Name string
	Age int
	friends []int64
}

func (u *User) HappyBirthday(){
	u.Age++
	fmt.Println("copy: ", u.Age)
}

func (u *User) Reset() {
	*u = User{}

	fmt.Println("Hello,", u.Name)
	fmt.Printf("inside: %p\n", u)
}


func (u *User) PrintName(){
	fmt.Println("Hello,", u.Name)
	fmt.Printf("inside: %p\n", u)
	
}

func main(){
	u := User{Age: 30, Name: "Alex"}
	u.PrintName()
	fmt.Printf("original: %p\n", &u)
	u.PrintName()
}