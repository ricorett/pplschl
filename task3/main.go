package main

import "fmt"

func main() {
	const (
		usdEur = 1.17
		usdRub = 75.19
	)
	//eurRub := usdRub * usdEur
	//fmt.Printf("%.2f\n ", eurRub)
	for {
		err := getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			err = getUserInput()
		}

	}
}

//func calcCurrency(num float64, firstConv float64, secondConv float64) (float64, float64) {
//
//}

func getUserInput() error {

	fmt.Print("Введите исходную валюту(USD, EUR, RUB): ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return err
	}

	var number float64
	fmt.Print("Введите сумму: ")
	_, err = fmt.Scanln(&number)
	if err != nil {
		return err
	}

	fmt.Print("Введите целевую валюту(USD, EUR, RUB): ")
	var inputDist string
	_, err = fmt.Scanln(&inputDist)
	if err != nil {
		return err
	}

	return nil
}
