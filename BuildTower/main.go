package main

import (
	"fmt"
	"strings"
)

func main() {
	nFloors := 2
	fmt.Println(TowerBuilder(nFloors))
}
func TowerBuilder(nFloors int) []string {
	var res []string
	var an strings.Builder

	for i := 1; i < nFloors+1; i++ {
		if nFloors == 1 {
			an.WriteString("*")
		} else {
			for j := i; j <= nFloors-1; j++ {
				an.WriteString(" ")
			}
			for k := 1; k < i*2; k++ {
				an.WriteString("*")
			}
			for j := i; j <= nFloors-1; j++ {
				an.WriteString(" ")
			}
		}
		
		res = append(res, an.String())
		an.Reset()

	}
	return res
}
