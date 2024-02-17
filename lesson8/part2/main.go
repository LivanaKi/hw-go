package main

import(
	"fmt"
	"errors"
)

type Checker struct{
	err error
	matched int
}

func (c *Checker) Step(s string){
	if c.err != nil {
		return
	}

	if len(s) == 0{
		c.err = errors.New("empty string")
		return
	}

	c.matched++
}

func (c *Checker) Error() error {
	return c.err
}

func main(){
	c := &Checker{}
	records := []string{"foo", "bar", "o", "test"}

	for _, r := range records{
		c.Step(r)
	}

	if err := c.Error(); err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("Ok: %d", c.matched)
}