package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-*!")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimestamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	pass := make([]rune, n)
	for i := range pass {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(pass)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimestamp, error) {

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("url parse error")
	}

	if login == "" {
		return nil, errors.New("login required")
	}

	newAcc := &AccountWithTimestamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
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
