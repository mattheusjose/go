package main

import "fmt"

func main() {

	numbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}

	mmc := minimumMultipleCommon(numbers)

	fmt.Printf("MMC is %v\n", mmc)

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

func nextPrimeFactor(factor int) int {

	var (
		primeSearchLimit int = 1_000_000
	)

	for index := factor + 1; index < factor*primeSearchLimit; index++ {

		if isPrimeNumber(index) {
			return index
		}
	}

	return factor
}

func existsAnyDivibleByFactor(numbers []int, factor int) bool {

	for _, number := range numbers {
		if number%factor == 0 {
			return true
		}
	}

	return false
}

func calculateMmc(numbers []int, factors []int) int {

	for _, number := range numbers {
		if number > 1 {
			return 0
		}
	}

	mmc := 1
	for _, factor := range factors {
		mmc *= factor
	}

	return mmc
}

func minimumMultipleCommon(numbers []int) int {

	var (
		mmc     int = 0
		factor  int = 2
		factors []int
	)

	for mmc == 0 {

		appendFactorToList := false

		for index := 0; index < len(numbers); index++ {

			if numbers[index]%factor == 0 {
				numbers[index] = numbers[index] / factor
				appendFactorToList = true
			}
		}

		if appendFactorToList {
			factors = append(factors, factor)
		}

		if !existsAnyDivibleByFactor(numbers, factor) {
			factor = nextPrimeFactor(factor)
		}

		mmc = calculateMmc(numbers, factors)
	}

	return mmc
}
