package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close() 

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "END" {
			return
		}
		n, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			fmt.Println("> That not number")
		} else {
			fmt.Println(">", n)
		}
	}
}
