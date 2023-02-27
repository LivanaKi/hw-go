package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(Zeros(30))
}
func rgb(r, g, b int) string { 
	var res strings.Builder 
	var an []int 
	
	for i := 0; i < 3; i++ { 
	 an = append(an, r, g, b) 
	
	 if an[i] > 255 { 
	  an[i] = 255 
	 } 
	 if an[i] < 0 { 
	  an[i] = 0 
	 } 
	 if an[i] < 12 { 
	  res.WriteString(strings.ToUpper(strconv.FormatInt(0, 16))) 
	  res.WriteString(strings.ToUpper(strconv.FormatInt(int64(an[i]), 16))) 
	 } else { 
	  res.WriteString(strings.ToUpper(strconv.FormatInt(int64(an[i]), 16))) 
	 } 
	} 
	
	return res.String() 
   }
func Zeros(n int) int {
	divisor := 5
	var result int
	for true {
		temp := n / divisor
		if temp < 1 {
			break
		} else {
			result = result + temp
			divisor = divisor * 5
		}
	}
	return result
}
