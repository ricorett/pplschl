package main

import (
	"fmt"
	// "os"
	"passwordManager/account"
	"passwordManager/encrypter"
	"passwordManager/output"

	"passwordManager/files"

	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu =  map[string]func(*account.VaultWithDb){
	"1" : createAccount,
	"2" : findAccountByUrl,
	"3" : findAccountByLogin,
	"4" : deleteAccount,
}

func main() {

	
	// vault := account.NewVaultWithDb(cloud.NewCloudDB("https://a.ru"))
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалость найти env файл")
	}
	vault := account.NewVaultWithDb(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())

Menu:
	for {
		choose := promptData([]string{
		"1. Создать аккаунт", 
		"2. Найти аккаунт по URL",
		"3. Поиск по логину",
		"4. Удалить аккаунт",
		"5. Выйти",
		"Выберите вариант",
		})
		menuFunc := menu[choose]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		
	}
}



func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите юрл для удаления"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Deleted")

	} else {
		output.PrintError("Не найдено")
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите юрл для поиска"})
	accounts := vault.FindAccounts(url, func(acc account.Account, str string)bool{
		return strings.Contains(acc.Url, str)
	} )
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин для поиска"})
	accounts := vault.FindAccounts(login, func(acc account.Account, str string)bool{
		return strings.Contains(acc.Login, str)
	} )
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account){
	if len(*accounts) == 0 {
		color.Red("Accounts not found")
	}
	for _, acc := range *accounts {
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
