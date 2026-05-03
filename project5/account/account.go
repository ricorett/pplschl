package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"

	"github.com/fatih/color"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!")

type Account struct {
	Login    string `json : "login"`
	Password string `json : "password"`
	Url      string `json : "url"`
}

//type AccountWithTimestamp struct {
//	Account
//	createdAt time.Time
//	updatedAt time.Time
//}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
	fmt.Println(acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	pass := make([]rune, n)
	for i := range pass {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(pass)
}

func NewAccount(login, password, urlString string) (*Account, error) {

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("url parse error")
	}

	if login == "" {
		return nil, errors.New("login required")
	}

	newAcc := &Account{
		Login:    login,
		Url:      urlString,
		Password: password,
	}

	if password == "" {
		newAcc.generatePassword(12)

	}
	return newAcc, nil
}
