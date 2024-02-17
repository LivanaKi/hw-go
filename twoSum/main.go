package main

import (
	"fmt"

	"github.com/Users/natza/hw-go/twoSum/sum"
)

func main(){
	var graph = map[string][]string{
		"you":    []string{"alice", "bob", "claire"},
		"bob":    []string{"alice","anuj", "peggym"},
		"alice":  []string{"peggym"},
		"claire": []string{"jonny", "tho"},
		"anuj":   []string{},
		"peggym":  []string{},
		"tho":   []string{},
		"jonny":  []string{},
	}

	fmt.Println(sum.Graph(graph))

/*
	for i := 1; i < 100; i++{
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
	fmt.Println()


	fmt.Println("Prime numbers less than 20:")

    for number := 1; number < 20; number++ {
        if sum.Findprimes(number) {
            fmt.Printf("%v ", number)
        }
    }
	fmt.Println()

	num := 1994
	fmt.Println(sum.IntegerRom(num))
	fmt.Println()

	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(sum.MaxArea(s))
	fmt.Println()

	firstName := "John"
	sum.UpdateName(&firstName)
	fmt.Println(firstName)
	fmt.Println()
	
	sum.Highlow(2, 0)
	fmt.Println("Program finished successfully!")*/
}
