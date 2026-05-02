package main

import (
	"fmt"
	"passwordManager/account"
)

func main() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")

	account1, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("неверный формат url")
		return
	}

	account1.OutputPassword()

}

func promptData(prompt string) string {

	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
