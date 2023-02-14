package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aA11"
	fmt.Println(duplicate_count(str))
}
func duplicate_count(s string) int {
	val := 0
	word := make(map[string]int)
	
	s1 := strings.ToUpper(s)
	
	for _, val := range s1 {
		_, ok := word[string(val)]
		if ok {
			word[string(val)] += 1
		} else {
			word[string(val)] = 1
		}
	}
	for _, value := range word{
		if value > 1{
			val++
		}
	}
	return val
}
