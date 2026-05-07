package config

import ("fmt"
	"os"
	"github.com/joho/godotenv"
)

type Env struct{
	Key string
}

func ReadEnv(){
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Не удалость найти env файл")
	}
	var e  Env
	e.Key = os.Getenv("KEY")
}
