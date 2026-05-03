package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/files"
)

func main() {
	//1 create account
	//2 find account
	//3 delete account
	//4 exit
	//createAccount()
Menu:
	for {
		choose := getMenu()
		switch choose {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			//deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var choose int

	fmt.Println("Выберите вариант")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выйти")
	_, err := fmt.Scan(&choose)
	if err != nil {
		return 0
	}

	return choose
}

func deleteAccount(login string) {

}

func findAccount() {}

func promptData(prompt string) string {

	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func createAccount() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")

	account1, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("неверный формат url")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*account1)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(data, "data.json")
}
