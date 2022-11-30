package main

import (
	"errors"
	"fmt"
)

func main() {

	var (
		n1 int
		n2 int
	)

	fmt.Printf("Type the first value: ")
	fmt.Scanf("%d\n", &n1)
	fmt.Printf("Type the second value: ")
	fmt.Scanf("%d", &n2)

	division, error := division(n1, n2)

	if error != nil {
		fmt.Println("Error: ", error)
	}

	fmt.Println(n1, "/", n2, "=", division)

	product, error := multiplication(n1, n2)
	if error != nil {
		fmt.Println("Error: ", error)
	}

	fmt.Println(n1, "x", n2, "=", product)
}

func division(divisor, dividend int) (int, error) {

	if dividend == 0 {
		return 0, errors.New("dividend must be greather than zero")
	}

	return divisor / dividend, nil
}

func multiplication(number, factor int) (int, error) {

	if number == 0 || factor == 0 {
		return 0, errors.New("either number and factor must be greather than zero")
	}

	return number * factor, nil
}
