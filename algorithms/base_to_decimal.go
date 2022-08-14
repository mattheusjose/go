package main

import (
	"fmt"
	"math"
	"strconv"
)

var digits = map[string]string{
	"A": "10",
	"B": "11",
	"C": "12",
	"D": "13",
	"E": "14",
	"F": "15",
}

func ToDecimal(numericRep string, base int) int {

	exponentialFactor := len(numericRep) - 1
	var decimalNumber int
	var digit string

	for _, bit := range numericRep {

		index := string(bit)
		mapDigit, hasKey := digits[index]

		digit = string(bit)

		if hasKey {
			digit = string(mapDigit)
		}

		numericDigit, err := strconv.Atoi(digit)
		if err != nil {
			panic("Error converting numeric value")
		}

		result := int(math.Pow(float64(base), float64(exponentialFactor))) * int(numericDigit)
		decimalNumber += result
		exponentialFactor -= 1
	}

	return decimalNumber
}

func main() {
	base := 2

	fmt.Println(fmt.Sprintf("%v to base %v = %v", "1", 2, ToDecimal("1", base)))
	fmt.Println(fmt.Sprintf("%v to base %v = %v", "10", 2, ToDecimal("10", base)))
	fmt.Println(fmt.Sprintf("%v to base %v = %v", "21", 3, ToDecimal("21", 3)))
	fmt.Println(fmt.Sprintf("%v to base %v = %v", "1110", 2, ToDecimal("1110", 2)))
	fmt.Println(fmt.Sprintf("%v to base %v = %v", "E", 16, ToDecimal("E", 16)))
	fmt.Println(fmt.Sprintf("%v to base %v = %v", "11", 16, ToDecimal("11", 16)))
}
