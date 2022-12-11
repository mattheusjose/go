package main

import "fmt"

func main() {

	lastPosition := 10001
	primes := getPrimeNumbers(lastPosition)

	fmt.Printf("Prime number at %vth position is %v", lastPosition, primes[lastPosition-1])
}

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

func getPrimeNumbers(numberOfElements int) []int {

	primes := []int{2}
	number := 3

	for len(primes) < numberOfElements {

		if isPrimeNumber(number) {
			primes = append(primes, number)
		}

		number += 2
	}

	return primes
}
