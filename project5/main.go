package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimestamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	pass := make([]rune, n)
	for i := range pass {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(pass)
}

//
//func newAccount(login, password, urlString string) (*account, error) {
//
//	_, err := url.ParseRequestURI(urlString)
//
//	if err != nil {
//		return nil, errors.New("url parse error")
//	}
//
//	if login == "" {
//		return nil, errors.New("login required")
//	}
//
//	newAcc := &account{
//		login:    login,
//		url:      urlString,
//		password: password,
//	}
//	if password == "" {
//		newAcc.generatePassword(12)
//
//	}
//	return newAcc, nil
//}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimestamp, error) {

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("url parse error")
	}

	if login == "" {
		return nil, errors.New("login required")
	}

	newAcc := &accountWithTimestamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			url:      urlString,
			password: password,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)

	}
	return newAcc, nil
}

func main() {

	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите url")

	account1, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("неверный формат url")
		return
	}

	account1.outputPassword()

}

func promtData(promt string) string {

	fmt.Print(promt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
