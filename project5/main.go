package main

import (
	"fmt"
	"passwordManager/account"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault()
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
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите юрл для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Deleted")

	} else {
		color.Red("Not Deleted")
	}
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите юрл для поиска")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Accounts not found")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func promptData(prompt string) string {

	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func createAccount(vault *account.Vault) {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")

	account1, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("неверный формат url")
		return
	}

	vault.AddAccount(*account1)

}
