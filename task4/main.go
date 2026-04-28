package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation := getInput()
	slice := getSlice()
	result := calcOperation(operation, slice)
	fmt.Printf("%s %d", operation, result)
}

func getInput() string {
	var operation string
	fmt.Print("Enter operation: ")
	fmt.Scan(&operation)

	return operation
}

func getSlice() []int {
	slice := make([]int, 0, 0)

	fmt.Print("Enter slice : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	parts := strings.Split(text, ",")

	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		slice = append(slice, n)

	}
	return slice
}

func calcOperation(operation string, slice []int) int {
	operation = strings.ToUpper(operation)
	result := 0
	switch operation {
	case "AVG":
		result = calcAvg(slice)
	case "SUM":
		result = calcSum(slice)
	case "MED":
		result = calcMed(slice)
	}
	return result
}

func calcAvg(slice []int) int {
	sum := 0
	for _, n := range slice {
		sum += n
	}
	avg := sum / len(slice)

	return avg
}

func calcSum(slice []int) int {
	sum := 0
	for _, n := range slice {
		sum += n
	}

	return sum
}

func calcMed(slice []int) int {
	var median int
	sort.Ints(slice)
	if len(slice)%2 == 0 {
		median = (slice[len(slice)/2-1] + slice[len(slice)/2]) / 2
	} else {
		median = slice[len(slice)/2]
	}
	return median
}
