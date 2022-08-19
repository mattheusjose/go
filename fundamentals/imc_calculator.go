package main

import (
	"fmt"
)

func womanImc(height float32) float32 {

	imc := (62.2 * height) - 44.7
	return imc
}

func manImc(height float32) float32 {

	imc := (72.7 * height) - 58.0
	return imc
}

func main() {

	handlers := map[string]func(float32) float32{
		"M": manImc,
		"F": womanImc,
	}

	var gender string
	var height float32

	fmt.Print("Type your gender: ")
	fmt.Scanf("%s\n", &gender)

	fmt.Print("Type your height: ")
	fmt.Scanf("%f", &height)

	handler := handlers[gender]
	imc := handler(height)

	fmt.Println("Gender: ", gender)
	fmt.Println("Height: ", height)
	fmt.Println("Imc: ", imc)
}
