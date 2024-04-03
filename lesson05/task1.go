package main

import(
	"fmt"
)
type IntStack struct{
	number []int
}
func (i *IntStack) Push(num int) {
	i.number = append(i.number, num)
}

func (i *IntStack) Pop() int{
	var num int
	if i.number != nil{
		num = i.number[len(i.number)-1]
		i.number = i.number[:len(i.number)-1]
	}else{
		return 0
	}
	return num
}
func main(){
	s := IntStack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("expected 30, got %d\n", s.Pop())
	fmt.Printf("expected 20, got %d\n", s.Pop())
	fmt.Printf("expected 10, got %d\n", s.Pop())

}