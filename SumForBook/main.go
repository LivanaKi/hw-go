package main

import "fmt"

type List struct {
	data int
	next *List
}

func main() {
	//list := &List{data: 3, next: &List{data: 4, next: &List{data: 3, next: nil}}} //answer 10
	//list := &List{data: 3, next: nil}  --- answer 3
	//list := &List{next: nil}  ---  answer 0
	//list := &List{} ---- answer 0
	//fmt.Println(Sum(list, 0))

	// ---- another solution ----- 

	list := []int{3, 4, 5, 3, 1, 7}
	//fmt.Print(sumForArr(list))
	//fmt.Print(lenForArr(list))
	fmt.Println(maxForArr(list))
}
func sumForArr(list []int) int{
	if len(list) == 0{
		return 0
	}
	return list[0] + sumForArr(list[1:])
}

func lenForArr(list []int) int{
	if len(list) == 0{
		return 0
	}
	return 1 + lenForArr(list[1:])
}

func maxForArr(list []int) int{
	if len(list) == 0{
		return 0
	}
	if list[0] < maxForArr(list[1:]){
		return maxForArr(list[1:])
	}else{
		return list[0]
	}
}

func Sum(list *List, sum int) int {

	if list.next == nil {
		sum += list.data
		return sum
	}
	sum += list.data
	fmt.Println(sum)
	return Sum(list.next, sum)
}
