package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(time.Format("2 Jan 2006 15:04:05"))
}
