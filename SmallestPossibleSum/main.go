package main

import (
	"fmt"
)

func main() {

	fmt.Println(Solution([]int{1, 12, 96, 48, 60, 24, 72, 36, 72, 72, 48}))

}

func Solution(ar []int) int {
	var res int
	for i, j := len(ar)-1, len(ar)-2; i >= 0 && j >= 0; i, j = i-1, j-1 {
        for ar[j] > ar[i]{
			ar[j] -= ar[i]
		}
    }

	for i := 0; i < len(ar); i++ {
		res += ar[i]
	}
	return res
}
