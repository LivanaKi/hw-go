package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	for i := 33; i >= 1; i-- {
		b.WriteString("Код")
		b.WriteRune('ь')
	}

	result := b.String()
	fmt.Println(result)

}
