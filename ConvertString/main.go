package main

import (
	"fmt"
	"strings"
)

func main() {
	strin := "the_stealth_warrior"
	fmt.Println(ToCamelCase(strin))
}
func ToCamelCase(s string) string {
	var res []string
	var res1 []string
	var an strings.Builder
	for _, val := range s {
		res = append(res, string(val))
	}

	for i := 0; i < len(res); i++ {
		if res[i] == "-" || res[i] == "_" && i+1 < len(res) {
			res[i] = ""
			res[i+1] = strings.ToUpper(res[i+1])
		} else {
			res1 = append(res1, res[i])
		}
	}
	for j := 0; j < len(res1); j++ {
		an.WriteString(res1[j])
	}
	return an.String()
}
