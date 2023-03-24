package main

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
	"unicode"
)
var ErrInvalidString = errors.New("invalid string")
func main() {
	start := "aaa10b"
	result, err := Unpack(start)
	if err == nil{
	fmt.Println("Unpacked string:", result)
	}else{
		fmt.Println("Unpacked string:", err)
	}
}

func Unpack(start string) (string, error) {

	var result strings.Builder

	if len(start) == 0 {
		return "", nil
	} else {
		for i, solution := range start {
			if unicode.IsNumber(rune((start[0]))) {
				return "", ErrInvalidString
			} else if unicode.IsDigit(solution) {
				num, _ := strconv.Atoi(string(start[i]))
				if num == 1 {
					form := start[i+1]
					if unicode.IsDigit(rune(form)) {
						num1, _ := strconv.Atoi(string(start[i+1]))
						if num1 == 0 {
							return "", ErrInvalidString
						}

					}
				}
				if num == 0 {
					return "", ErrInvalidString
				}
				result.WriteString(strings.Repeat(string(start[i-1]), num-1))

			}
			if unicode.IsLetter(solution) {
				result.WriteRune(solution)
			}
		}
	}
	return result.String(), nil
}
