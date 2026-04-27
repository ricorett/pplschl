package main

import (
	"fmt"
	"strings"
)

const (
	usdEur = 1.17
	usdRub = 75.19
	eurRub = usdRub * usdEur
)

func main() {

	//fmt.Printf("%.2f\n ", eurRub)

	var inputSrc, inputDist string
	var number float64
	var err error
	for {
		inputSrc, number, inputDist, err = getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			inputSrc, number, inputDist, err = getUserInput()
		} else {
			break
		}
	}
	calc := calcCurrency(inputSrc, number, inputDist)
	fmt.Printf("%.1f %s", calc, inputDist)

}
func calcCurrency(inputSrc string, number float64, inputDist string) float64 {
	var calculation float64
	switch inputDist {
	case "eur":
		switch inputSrc {
		case "usd":
			calculation = number * usdEur
		case "rub":
			calculation = number * usdRub
		}
	case "usd":
		switch inputSrc {
		case "eur":
			calculation = number * usdEur
		case "rub":
			calculation = number * usdRub
		}

	case "rub":
		switch inputSrc {
		case "eur":
			calculation = number * eurRub
		case "usd":
			calculation = number * usdRub
		}
	default:
		fmt.Println("Error")
	}

	return calculation
}

func getUserInput() (string, float64, string, error) {

	fmt.Print("Введите исходную валюту(USD, EUR, RUB): ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", 0, "", err
	}
	input = strings.ToLower(input)

	var number float64
	fmt.Print("Введите сумму: ")
	_, err = fmt.Scanln(&number)
	if err != nil {
		return "", 0, "", err
	}

	fmt.Print("Введите целевую валюту(USD, EUR, RUB): ")
	var inputDist string
	_, err = fmt.Scanln(&inputDist)
	if err != nil {
		return "", 0, "", err
	}
	inputDist = strings.ToLower(inputDist)

	return input, number, inputDist, nil
}
