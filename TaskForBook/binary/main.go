package main

import (
	"fmt"
)

func main(){
	fmt.Println(binarySearch([]string{"angel", "bar", "live", "zero"}, "bar"))
}

func binarySearch(list []string, n string) bool{
	first := 0
	last := len(list) - 1

	for last >= first{
		mid := last + first / 2

		if list[mid] == n{
			return true
		}else{
			if n < list[mid]{
				last = mid - 1
			}else{
				first = mid + 1
			}
		}
	}
	return false
}