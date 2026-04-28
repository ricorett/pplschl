package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const IMTPower = 2

func main() {
	for {
		err := calcIMT()
		if err != nil {
			fmt.Printf("Error calculating IMT: %v\n", err)
		}
		choose := retryInput()
		if choose {
			continue
		} else {
			break
		}

	}
}

func getUserInput() (userHeight float64, userKg float64) {

	fmt.Println("___ Калькулятор индекса массы тела___")
	fmt.Print("Введите рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес: ")
	fmt.Scan(&userKg)

	return
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массыл тела: %.1f \n", imt)
	fmt.Print(result)
}

func calcIMT() error {

	userHeight, userKg := getUserInput()
	if userHeight <= 0 || userKg <= 0 {
		err := errors.New("invalid input")
		return err
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	outputResult(IMT)
	return nil
}

func retryInput() bool {
	cont := false
	fmt.Println("Повторить ввод? ")
	var answ string
	fmt.Scan(&answ)
	answ = strings.ToLower(answ)
	if answ == "y" {
		cont = true
	} else {
		cont = false
	}
	return cont
}
