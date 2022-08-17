package main

import (
	"fmt"
)

func SquareArea(base int) int {

	area := base * base
	return area
}

func main() {

	var squareBase int
	fmt.Print("Base of square in cm: ")
	fmt.Scanf("%d\n", &squareBase)

	squareArea := SquareArea(squareBase)

	fmt.Println(fmt.Sprintf("Square with base %vcm has %vcm²", squareBase, squareArea))

}
