package main

import (
	"fmt"
)

func main() {
	n := 5
	fmt.Println(Perimeter(n))
}
func Perimeter(n int) int {
	var sum int 
	sl := make([]int, n+1)
		for i, _ := range sl{
			if i <= 1{
				sl[i] = 1
				sum += 4* sl[i]
			}else{
					sl[i] = sl[i-1] + sl[i-2]
					sum += 4* sl[i]
			}
		}
	return sum
}