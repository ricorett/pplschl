package files

import (
	"fmt"
	"os"
	"passwordManager/output"
)

type JsonDB struct {
	filename string
}



func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

func (db *JsonDB) Read() ([]byte, error) {

	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	return data, nil
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}

	fmt.Println("File written to: " + db.filename)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			output.PrintError(err)
		}
	}(file)
}
