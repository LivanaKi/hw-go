package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats. ")
		os.Exit(1)
	}

	argument := os.Args
	var err error = errors.New("An error")
	k := 1
	
	for err != nil {
		if k >= len(argument) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		_, err = strconv.ParseFloat(argument[k], 64)
		k++
	}
	var sum float64
	for i := 1; i < len(argument); i++ {
		n, err := strconv.ParseFloat(argument[i], 64)
		if err == nil {
			sum += n
		}
	}

	fmt.Println("Sum: ", sum)
}