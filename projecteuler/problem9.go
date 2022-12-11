package main

import "fmt"

func main() {

	target := 1000
	triplet := pythagoreanTriplet(target)
	product := calculateProduct(triplet)
	fmt.Printf("Triplet is %v\n", triplet)
	fmt.Printf("Product is %v\n", product)

}

func calculateProduct(numbers []int) int {

	product := 1
	for _, number := range numbers {
		product *= number
	}
	return product
}

func pow(number, base int) int {

	product := 1
	for index := 0; index < base; index++ {
		product *= number
	}

	return product
}

func isPerfectSquare(a, b, c int) bool {

	squareA := pow(a, 2)
	squareB := pow(b, 2)
	squareC := pow(c, 2)

	return squareA+squareB == squareC
}

func pythagoreanTriplet(target int) []int {

	for a := 1; a <= target/3; a++ {
		for b := 1; b < target/2; b++ {
			for c := 1; c < target; c++ {

				if a+b+c == target && isPerfectSquare(a, b, c) {
					return []int{a, b, c}
				}
			}
		}
	}

	return []int{}
}
