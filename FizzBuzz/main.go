package main

import(
	"fmt"
)

func main(){
for i := 1; i <= 16; i++{
	switch i != 0 {
	case i%3 == 0 && i%5 == 0:
		fmt.Println("FizzBuzz")
	case i%3 == 0:
		fmt.Println("Fizz")
	case i%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(i)
	}
}
}