package main

import(
	//"fmt"
	"io/ioutil"
)

type PathError struct {
	Op string 
	Path string
	Err error
}

func (e *PathError) Error() string{
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func chekConfig(path string) error{
	_, err := ioutil.ReadFile(path)
	if err != nil{
		return &PathError{"open", path, err}
	}
	return nil
}

/*func main() {
	err := chekConfig("/etc/apt/sources.list")
	switch err := err.(type) {
	case *PathError:
		fmt.Println("path error: ", err)
	case nil:
		fmt.Println("success")
	default:
		fmt.Println("unknown error")
	}
}*/