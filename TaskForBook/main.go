package main

import(
	"fmt"
)

func main(){
	printNum(1)
}

func printNum(n int) int{
	if n > 10 {
		return 1
	}
	fmt.Println(n)
	return printNum(n + 1)
}