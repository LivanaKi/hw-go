package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)


func main() {
	n1 := "3715290469715693021198967285016729344580685479654510946723"
	n2 := "68819615221552997273737174557165657483427362207517952651"
	fmt.Println(LastDigit(n1, n2))
}

func LastDigit(n1, n2 string) int {
	var re = regexp.MustCompile(`[[:punct:]]`)
	n1 = re.ReplaceAllString(n1, "")
	n2 = re.ReplaceAllString(n2, "")
	
	if n2 == "0" {
		return 1
	}
	num1 := 0
	if len(n2) > 2 {
		for i := 2; i > 0; i-- {
			num1 = num1*10 + int(rune(n2[len(n2)-i])-'0')
		}
	}else{
		num1, _ = strconv.Atoi(n2)
	}
	if n2 == "0" {
		return 1
	}
	num, _ := strconv.Atoi(string(n1[len(n1)-1]))
	res := num1 % 4
	result := 0
	if res == 0{
		result = int(math.Pow(float64(num), float64(4)))
	}else{
		result = int(math.Pow(float64(num), float64(res)))
	}
	return result%10
}
