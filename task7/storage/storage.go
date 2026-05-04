package storage

import (
	_ "bins/bin"
	"encoding/json"
	"os"
)

func SaveBin(BinList List, filename string) error {
	file, err := json.Marshal(BinList)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

func ReadBin(BinList List, filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &BinList)
	if err != nil {
		return err
	}
	return nil
}
