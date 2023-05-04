package main

import (
	"errors"
	//"fmt"
)

type Cheker struct{
	err error
	matched int
}

func (c *Cheker) Step(s string){
	if c.err != nil{
		return
	}
	if len(s) == 0{
		c.err = errors.New("empty string")
		return
	}
	c.matched++
}
func (c *Cheker) Errorf() error{
	return c.err
}
/*func main(){
	c := &Cheker{}
	records := []string{"foo", "bar", "test"}
	
	for _, r := range records{
		c.Step(r)
	}
	if err := c.Errorf(); err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("OK: %d", c.matched)
}*/