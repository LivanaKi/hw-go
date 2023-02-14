package main

import (
	"fmt"
)

func main() {
	size := 3
	fmt.Println(MultiplicationTable(size))
}
func MultiplicationTable(size int) [][]int {

	var res [][]int

	for i := 0; i < size; i++ {
		res = append(res, []int{})
		for j := 0; j < size; j++ {
			res[i] = append(res[i], 0)
			for k := 1; k <= i+1; k++ {
				for m := 1; m <= j+1; m++ {
					//res[i][j] = k
					res[i][j] = k * m
				}
			}
		}
	}
	return res
}
