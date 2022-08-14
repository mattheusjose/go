package main

import (
	"fmt"
)

/*

base 2 = max digit is 1
base 3 = max digit is 2
base 4 = max digit is 3
...
base 16 max digit is:
A = 10
B = 11
C = 12
D = 13
E = 14
F = 15
*/

var digits = map[string]string{
	"10": "A",
	"11": "B",
	"12": "C",
	"13": "D",
	"14": "E",
	"15": "F",
}

func DecimalToAnyBase(decimal int, base int) string {

	number := decimal
	numericRepresentation := ""

	for number > 0 {

		remainder := number % base
		digit := fmt.Sprintf("%v", remainder)

		if remainder >= 10 {
			digit = digits[digit]
		}

		numericRepresentation = digit + numericRepresentation
		number = number / base
	}

	return numericRepresentation
}

func main() {

	base := 16
	decimal := 17

	fmt.Println(fmt.Sprintf("%v base %v = %v", decimal, base, DecimalToAnyBase(decimal, base)))
}
