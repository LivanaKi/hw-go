package main

import (
	//"fmt"

//	"errors"
)

/*const testEmail = "test@gmail.com"
//Что б проверить, соответстует ли ошибка типу ее надо развернуть

// создали свою ошибку
type UserNotFound struct{

}

func (e *UserNotFound) Error() string{
	return "user not found"
}

func getUserEmail(id int) (string, error){
	if id == 1{
		return testEmail, nil
	}
	return "", &UserNotFound{}
}

func loginUserById(id int)(bool, error){
	email, err := getUserEmail(id)
	if err != nil {
		return false, fmt.Errorf("unable to get email: %w", err) // другой варіант Wrap(). %w в старих версий нет, только начиная с 1,13
	}
	if email != testEmail{
		return false, fmt.Errorf("unable to login user")
	}
	return true, nil
}
// схожий пример
func main() {
	_, err := loginUserById(1)

	switch err := errors.Unwrap(err).(type){ //разпаковка ошибкі с помощью стандартной бібліотекі errors
//только вот в errors.Cause(err).(type) в прошлом варіанте просто печатали.
//а в даном случае ми проверяем. Визиваем метод Cause() и разворачиваем 
//ошибку, что б увидеть полную картину ошибки, до тех пор пока не достанет 
//последнюю, если ошика nil, то вернется nil
	case *UserNotFound:
		fmt.Println("unable to find user: ", err)
	case nil:
		fmt.Println("success")
	default:
		fmt.Println("unknown error")
	}
}*/