//если не использивать сторонние пакети
package main

import(
	"fmt"
	"errors"
)

func foo() (bool, error) {
	return false, errors.New("foo failed")//исскуствино созданная ошибка
}

func bar() (bool, error) {
	return false, fmt.Errorf("user %q not found", "test")//исскуствино созданная ошибка, но более форматированная
}
//%q - индификатор, "test" - аргумент

func main(){
	_, err := foo()
	if err != nil{
		fmt.Println(err)
	}
	if _, err := bar(); err != nil{
		fmt.Println(err)
	}
}