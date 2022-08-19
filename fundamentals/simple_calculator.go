package main

import (
	"fmt"
)

func pow(number float32, power int) float32 {

	if power == 0 {
		return 1
	}

	return number * pow(number, power-1)
}

func main() {

	var integer1, integer2 int
	var decimal1 float32

	fmt.Print("Type the first integer: ")
	fmt.Scanf("%d\n", &integer1)

	fmt.Print("Type the second integer: ")
	fmt.Scanf("%d\n", &integer2)

	fmt.Print("Type the first decimal: ")
	fmt.Scanf("%f", &decimal1)

	result1 := integer1 * 2 * integer2 / 2
	result2 := float32(integer1*3) + decimal1
	result3 := pow(decimal1, 3)

	fmt.Println("result 1: ", result1)
	fmt.Println("result 2: ", result2)
	fmt.Println("result 3: ", result3)
}
