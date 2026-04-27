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
		}
		calc := calcCurrency(inputSrc, number, inputDist)
		fmt.Printf("%.1f %s \n", calc, inputDist)
	}

}
func calcCurrency(inputSrc string, number float64, inputDist string) float64 {
	var calculation float64
	switch inputSrc {
	case "eur":
		switch inputDist {
		case "usd":
			calculation = number / usdEur
		case "rub":
			calculation = number * eurRub
		case "eur":
			calculation = number
		}
	case "usd":
		switch inputDist {
		case "eur":
			calculation = number / usdEur
		case "rub":
			calculation = number * usdRub
		case "usd":
			calculation = number
		}

	case "rub":
		switch inputDist {
		case "eur":
			calculation = number / eurRub
		case "usd":
			calculation = number / usdRub
		case "rub":
			calculation = number
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

	if !validateCurrency(input) {
		return "", 0, "", fmt.Errorf("invalid input %s", input)
	}

	var number float64
	fmt.Print("Введите сумму: ")
	_, err = fmt.Scanln(&number)
	if err != nil {
		return "", 0, "", err
	}
	if number < 0 {
		return "", 0, "", fmt.Errorf("number must not be negative: %f", number)
	}

	fmt.Print("Введите целевую валюту(USD, EUR, RUB): ")
	var inputDist string
	_, err = fmt.Scanln(&inputDist)
	if err != nil {
		return "", 0, "", err
	}
	inputDist = strings.ToLower(inputDist)

	if !validateCurrency(inputDist) {
		return "", 0, "", fmt.Errorf("invalid input %s", inputDist)
	}

	return input, number, inputDist, nil
}

func validateCurrency(input string) bool {
	return input == "eur" || input == "rub" || input == "usd"
}
