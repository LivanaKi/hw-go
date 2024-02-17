package main

import(
	"fmt"
	"io/ioutil"
)

type PathError struct{
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string{
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func checkConfig(path string) error{
	_, err := ioutil.ReadFile(path)
	if err != nil{
		return &PathError{"oprn", path, err}
	}
	return nil
} 

func main(){
	err := checkConfig("/etc/apt/source.list")
	switch err := err.(type){
	case *PathError:
		fmt.Println("path error: ", err)
	case nil:
		fmt.Println("success")
	default:
		fmt.Println("unknown error")
	}
}