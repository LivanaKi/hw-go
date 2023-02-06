package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "middle"

	

	var result strings.Builder
	var length int

	for i, val := range s {
		length = len(s)
		if length%2 == 0 {
			length /= 2
			if i == length-1 || i == length {
				result.WriteRune(val)
			}
		} else {
			length -= 1
			length /= 2
			if i == length {
				result.WriteRune(val)
			}
		}
	}

	//fmt.Println(length)
	fmt.Print(result.String())
}
