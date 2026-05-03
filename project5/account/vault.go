package account

import (
	"encoding/json"
	"time"
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
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (v *Vault) AddAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
}
