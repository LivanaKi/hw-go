package main

import (
	"fmt"
)

func itoa(i int) (s string) {
	var current int
	var negative int

	if i < 0{
		i *= -1
		negative = -1
	}

	if i == 0 {
		return "0"
	}

	for i != 0 {
		current = i % 10
		i /= 10
		
		s = string('0' + current) + s
	}

	if negative == -1{
		s = "-" + s
	}
	
	return s
}

func main() {

	type pair struct {
		i int
		s string
	}
	test := []pair{
		{0, "0"},
		{22, "22"},
		{32432523, "32432523"},
		{-35, "-35"},
	}
	for _, t := range test {
		if t.s == itoa(t.i) {
			fmt.Printf("#{t.i} #{t.s} - True")
		} else {
			fmt.Printf("#{t.i} #{t.s} - False")
		}
	}
	//a := "♬"
	//a := "Jack"
	//a := "Джек"

	//fmt.Println(len(a))

	//for id, val := range a {
	//	fmt.Printf("id %d %s\n", id, string(val))
	//}
}
