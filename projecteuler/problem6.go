package main

import "fmt"

func buildSquareList(numberOfElements int) []int {

	numbers := []int{}

	for index := 1; index <= numberOfElements; index++ {
		numbers = append(numbers, index)
	}

	return numbers
}

func sumOfSquare(numbers []int) int {

	squareSum := 0

	for _, number := range numbers {

		squareSum += number * number
	}

	return squareSum
}

func squareOfTheSum(numbers []int) int {

	squareSum := 0
	for _, number := range numbers {

		squareSum += number
	}

	return squareSum * squareSum
}

func main() {

	numberOfElements := 100
	numbers := buildSquareList(numberOfElements)

	sumOfSquare := sumOfSquare(numbers)
	squareSum := squareOfTheSum(numbers)
	diff := squareSum - sumOfSquare

	fmt.Printf("Sum of square %v is %v\n", numbers, sumOfSquare)
	fmt.Printf("Square of the sum %v is %v\n", numbers, squareSum)
	fmt.Printf("%v-%v=%v\n", squareSum, sumOfSquare, diff)

}
