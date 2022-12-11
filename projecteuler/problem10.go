package main

import "fmt"

func isPrimeNumber(number int) bool {

	if number == 1 {
		return false
	}

	index := 2
	for index*index <= number {

		if number%index == 0 {
			return false
		}

		index += 1
	}

	return true
}

func main() {

	sumOfPrimes := 2
	limit := 2_000_000

	for index := 3; index < limit; index += 2 {

		if isPrimeNumber(index) {
			sumOfPrimes += index
		}
	}

	fmt.Printf("The sum of primes bellow %v is %v", limit, sumOfPrimes)
}
