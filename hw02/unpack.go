package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var start string
	fmt.Println("Please enter string: ")

	fmt.Fscan(os.Stdin, &start)
	//start := ""
	fmt.Println("Unpacked string: ", Unpack(start))
}

func Unpack(start string) string {
	var result strings.Builder
	if len(start) == 0 {
		result.WriteString("")
	} else {
		for i, solution := range start {
			if unicode.IsNumber(rune((start[0]))) {
				result.Reset()
				result.WriteString("Don't correct")
				break
			}

			if unicode.IsDigit(solution) {
				num, _ := strconv.Atoi(string(start[i]))
				if num == 1 {
					form := start[i+1]
					if unicode.IsDigit(rune(form)) {
						num1, _ := strconv.Atoi(string(start[i+1]))
						if num1 == 0 {
							result.Reset()
							result.WriteString("Don't correct")
							break
						}

					}
				}
				if num == 0 {
					result.Reset()
					result.WriteString("Don't correct")
					break
				}
				result.WriteString(strings.Repeat(string(start[i-1]), num-1))

			}
			if unicode.IsLetter(solution) {
				result.WriteRune(solution)
			}
		}
	}
	return result.String()
}