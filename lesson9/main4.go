package main

import(
	"fmt"
)

func PrintAll(vals interface{}){
	for _, val := range vals.([]string){
		fmt.Println(val)
	}
}

/*func main(){
	names := []string{"stanley", "davsd", "oscar"}
	PrintAll(names)
}*/