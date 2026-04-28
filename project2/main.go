package main

import (
	"fmt"
)

func main() {
	transactions := make([]float64, 0)
	for {
		n := getInput()
		if n == 0 {
			break
		}
		transactions = append(transactions, n)
	}
	fmt.Println("Transactions: ", transactions)
	balance := calcTransactions(transactions)
	fmt.Println("Balance: ", balance)
}

func getInput() float64 {
	var n float64
	fmt.Print("Enter transaction (n for exit): ")
	fmt.Scan(&n)
	return n
}

func calcTransactions(transactions []float64) float64 {
	balance := 0.0
	for i := range transactions {
		balance += transactions[i]
	}
	return balance
}
