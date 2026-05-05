package account

import (
	"encoding/json"
	"passwordManager/output"
	"strings"
	"time"

	
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct{
	Vault
	db Db
}

type ByteReader interface{
	Read() ([]byte, error)
}


type ByteWriter interface{
	Write([]byte)
}

type Db interface{
	ByteWriter
	ByteReader
}

func (v *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func NewVaultWithDb(db Db) *VaultWithDb {
	
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault : Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
			}, 
			db : db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError(err)
	}
	return &VaultWithDb{
		Vault: vault,
		db: db,
	}
}

func (v *VaultWithDb) AddAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.save()
}

func (v *VaultWithDb) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, acc := range v.Accounts {
		isMatches := strings.Contains(url, acc.Url)
		if isMatches {
			accounts = append(accounts, acc)
		}

	}
	return accounts
}

func (v *VaultWithDb) DeleteAccountByUrl(url string) bool {
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
	v.save()
	return isDeleted
}

func (vault *VaultWithDb) save(){
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError(err)
	}
	vault.db.Write(data)

}