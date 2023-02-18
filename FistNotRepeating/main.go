package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "stress"

	fmt.Println(FirstNonRepeating(str))
}
func FirstNonRepeating(str string) string {
	var res string
	var min int
	var ans []rune

	lit := make(map[string]int)
	if len(str) == 0 {
		res = ""
	} else if len(str) == 1 {
		res = str
	} else {
		for i, vali := range str {
			ans = append(ans, vali)
			val := strings.ToLower(string(vali))
			_, ok := lit[string(val)]
			if ok {
				lit[string(val)] = 0
			} else {
				lit[string(val)] = i+1
			}
		}
		
		values := []int{}
		for _, value := range lit {
			if value != 0 {
				values = append(values, value)
			}
		}
		
		if len(values) == 0 {
			res = ""
		} else {
			min = values[0]
			for i := 1; i < len(values); i++ {
				if min > values[i] {
					min = values[i]
				}
			}

			for i, val := range ans {
				if i == min-1 {
					res = string(val)
				}
			}
		}
}
	return res
}
