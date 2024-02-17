package main

import (
	"fmt"
	//"strconv"
)

func main(){
	var page uint64
	page = 400000000000000000
	fmt.Println(PageDigits(page))
}
func PageDigits(pages uint64) uint64 {
    var result uint64
	result = 9
	for i := 2; i < 15; i++{
		result = (result*10)*uint64(i)
		fmt.Println(result)
	}	
	fmt.Println()
	fmt.Println(result)
	result = (result - pages)*11
	fmt.Println(7088888888888888907)
	return result
}