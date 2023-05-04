package main

import (
	//"fmt"
)

/*const testEmail = "test@gmail.com"

func getUserEmail(id int) (string, error){
	if id == 1{
		return testEmail, nil
	}
	return "", fmt.Errorf("email not found")//создали обичную ошибку
}

func loginUserById(id int) (bool, error){
	email, err := getUserEmail(id)
	if err != nil{
		return false, fmt.Errorf("unable to get email: %w", err)// в етой часте заврапили (wrap)
	}
	if email != testEmail{
		return false, fmt.Errorf("unable to login user")
	}
	return true, nil
}

func main(){
	if _, err := loginUserById(0); err != nil{
		fmt.Printf("%+v\n", err)//здесь печатаем ошибку
		// "%+v\n" ето разширений вариант и в даном коде ничего не виведится
		//а вот в pkg/errors вивелось, и в етом его преимущество
	}else{
		fmt.Println("success")
	}
}*/