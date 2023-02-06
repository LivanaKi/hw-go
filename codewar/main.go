package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "This is another test"
	fmt.Print(SpinWords(str))
}

func SpinWords(str string) string {
	var answear string
	var four strings.Builder
	var res []string

	one := strings.Split(str, " ")
	for i := 0; i < len(one); i++ {
		strs := one[i]
		length := len(strs)
		if length >= 5 {
			var two []string
			var three []string
			for _, val := range strs {

				two = append(two, string(val))
			}
			for j := len(two) - 1; j >= 0; j-- {
				three = append(three, two[j])
			}
			for k := 0; k < len(three); k++ {
				four.WriteString(three[k])
			}
			result := four.String()
			four.Reset()

			res = append(res, result)

		} else {
			res = append(res, strs)
		}
	}
	answear = strings.Join(res, " ")
	return answear
}
