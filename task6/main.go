package main

import (
	"fmt"
	"strings"
)

type Converter map[string]map[string]float64

func main() {
	m := Converter{
		"eur": {
			"usd": 1.17,
			"rub": 87.97,
		},
		"usd": {
			"eur": 1 / 1.17,
			"rub": 75.19,
		},
		"rub": {
			"usd": 1 / 75.19,
			"eur": 1 / 87.97,
		},
	}

	var inputSrc, inputDist string
	var number float64
	var err error
	for {
		inputSrc, number, inputDist, err = getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			inputSrc, number, inputDist, err = getUserInput()
		}
		calc := calcCurrency(inputSrc, number, inputDist, &m)
		fmt.Printf("%.1f %s \n", calc, inputDist)
	}

}
func calcCurrency(inputSrc string, number float64, inputDist string, m *Converter) float64 {
	return number * (*m)[inputSrc][inputDist]
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
