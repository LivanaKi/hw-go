package main

import(
	"fmt"
)

type Hound interface{
	Hunt()
}

type Poodle interface{
	Bark()
}
type Dog1 interface{
	Hound
	Poodle 
}

type GoldenRetriever struct{
	name string
}

func (GoldenRetriever) Hunt(){
	fmt.Println("hunt")
}
func (GoldenRetriever) Bark(){
	fmt.Println("bark")
}
/*
func f1(i Hound){
	i.Hunt()
}
func f2(i Poodle){
	i.Bark()
}

func main(){
	t := GoldenRetriever{"jack"}
	f1(t)
	f2(t)
}*/ 
//Даний пример можна немного сократить благодаря композиции

func f(i Dog1){i.Hunt(); i.Bark()}

/*func main(){
	t := GoldenRetriever{"jack"}
	f(t)
}*/