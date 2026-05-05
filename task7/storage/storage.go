package storage

import (
	"bins/bin"
	"encoding/json"
	"os"
)

type BinOp interface{
	SaveBin(bin.List, string) error
	ReadBin(*bin.List, string) error
}


func SaveBin(BinList bin.List, filename string) error {
	file, err := json.Marshal(BinList)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

func ReadBin(BinList *bin.List, filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, BinList)
	if err != nil {
		return err
	}
	return nil
}
