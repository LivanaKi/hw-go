package main

import (
	"fmt"
	"github.com/Users/natza/hw-go/lesson04/user"
)

func main() {
	a := user.User{
		Name: "Jack",
	}

	a.SetPass("pass")
	fmt.Printf("%v", a)
}