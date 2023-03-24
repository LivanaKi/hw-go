package main

import (
	"fmt"
)

type IntStack struct {
	arrArr []int
}

func (i *IntStack) Push(s int) {
	i.arrArr = append(i.arrArr, s)
}

func (i *IntStack) Pop() int {
	res := i.arrArr[len(i.arrArr)-1]
	i.arrArr = i.arrArr[:len(i.arrArr)-1]
	return res
}

func main() {
	s := IntStack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("epected 30, got %d\n", s.Pop())
	fmt.Printf("epected 20, got %d\n", s.Pop())
	fmt.Printf("epected 10, got %d\n", s.Pop())
}
