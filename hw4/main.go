package main

import(
	"fmt"
)

type List interface {
	Len() int
	/*Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)*/
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	listEl *ListItem
	len int
}

func (l *list) Len() int{
	return l.len
}

func main(){
	
}