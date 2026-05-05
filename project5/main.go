package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/output"

	"passwordManager/files"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVaultWithDb(files.NewJsonDB("data.json"))
	// vault := account.NewVaultWithDb(cloud.NewCloudDB("https://a.ru"))
Menu:
	for {
		choose := promptData([]string{
		"1. Создать аккаунт", 
		"2. Найти аккаунт",
		"3. Удалить аккаунт",
		"4. Выйти",
		"Выберите вариант",
		})
		switch choose {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

// func getMenu() int {
// 	var choose int

// 	fmt.Println("Выберите вариант")
// 	fmt.Println("")
// 	fmt.Println("2. Найти аккаунт")
// 	fmt.Println()
// 	fmt.Println("4. Выйти")
// 	_, err := fmt.Scan(&choose)
// 	if err != nil {
// 		return 0
// 	}

// 	return choose
// }

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите юрл для удаления"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Deleted")

	} else {
		output.PrintError("Не найдено")
	}
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите юрл для поиска"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Accounts not found")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt{
		if i == len(prompt) - 1 {
				fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	
	var res string
	fmt.Scanln(&res)
	return res
}

func createAccount(vault *account.VaultWithDb) {

	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите url"})

	account1, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Login")
		return
	}

	vault.AddAccount(*account1)

}
