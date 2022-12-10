package main

import (
	"fmt"
)

func PrimeFactors(number int) (primeFactors []int) {

	for number%2 == 0 {
		primeFactors = append(primeFactors, 2)
		number = number / 2
	}

	for index := 3; index*index <= number; index += 2 {
		for number%index == 0 {
			primeFactors = append(primeFactors, index)
			number = number / index
		}
	}

	if number > 2 {
		primeFactors = append(primeFactors, number)
	}

	return
}

func main() {

	var (
		numberInCheck int = 600851475143
	)

	fmt.Println(PrimeFactors(numberInCheck))
}
