package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func fibonacciSumOfEvenValues(limit int) int {

	var (
		sumOfEvenValues int = 0
		term1           int = 0
		term2           int = 1
		term3           int = 0
	)

	for index := 1; term3 < limit; index++ {

		term3 = term1 + term2
		term1 = term2
		term2 = term3

		if isEven(term3) {
			sumOfEvenValues += term3
		}
	}

	return sumOfEvenValues
}

func main() {

	var (
		limit int = 4_000_000
	)

	sumOfEvenValues := fibonacciSumOfEvenValues(limit)
	fmt.Printf("Sum: %v", sumOfEvenValues)
}
