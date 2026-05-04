package account

import (
	"encoding/json"
	"passwordManager/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (v *Vault) AddAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.UpdatedAt = time.Now()
	data, err := v.ToBytes()
	if err != nil {
		color.Red(" Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}

func (v *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, acc := range v.Accounts {
		isMatches := strings.Contains(url, acc.Url)
		if isMatches {
			accounts = append(accounts, acc)
		}

	}
	return accounts
}

func (v *Vault) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, acc := range v.Accounts {
		if acc.Url != url {
			accounts = append(accounts, acc)
		} else {
			isDeleted = true
		}

	}
	v.Accounts = accounts
	v.UpdatedAt = time.Now()
	data, err := v.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "data.json")

	return isDeleted
}
