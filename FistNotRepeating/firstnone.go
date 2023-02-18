package main

import "strings"

func firstNonRepeating(str string) string {
	var res string

	val := strings.ToLower(str)
	lit := make(map[rune]int)
	if len(str) == 0 {
		res = ""
	} else if len(str) == 1 {
		res = str
	} else {
		for _, vali := range val {
			lit[vali]++
		}
		for i, value := range val{
			if lit[value] == 1{
				return string(str[i])
			}

		}
	}
	return res}
