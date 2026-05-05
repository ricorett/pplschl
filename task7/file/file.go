package file

import (
	"os"
	"path/filepath"
)

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func IsJson(filename string) bool {
	ext := filepath.Ext(filename)
	if ext == ".json" {
		return true
	}
	return false
}
