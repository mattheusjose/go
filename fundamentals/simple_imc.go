package main

import (
	"fmt"
)

func calculateImc(height float32) float32 {

	imc := (72.7 * height) - 58.0
	return imc
}

func main() {

	var height float32

	fmt.Print("Type your height: ")
	fmt.Scanf("%f\n", &height)

	imc := calculateImc(height)

	fmt.Println(fmt.Sprintf("height: %vm", height))
	fmt.Println(fmt.Sprintf("Imc: %v", imc))

}
