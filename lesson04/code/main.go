package main

import (
	"fmt"
	//"strconv"
)

func main() {
	number := 130
	fmt.Println(itoa(number))

}

func itoa(num int) string {
	var ans string
	var negative = ""
	
	if num == 0{
		ans = "0"
	}
	if num < 0 {
		negative = "-"
		num = 0 - num
	}
	for num > 0 {
		num1 := num % 10
		num /= 10
		ans = string('0'+num1) + ans
	}
	return negative + ans
}
