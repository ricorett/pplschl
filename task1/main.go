package main

import "fmt"

func main() {
	const (
		usdEur = 1.17
		usdRub = 75.19
	)
	eurRub := usdRub * usdEur
	fmt.Printf("%.2f\n", eurRub)
}
