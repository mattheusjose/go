package main

import (
	"fmt"
)

func main() {

	var fahrenheit float32
	fmt.Print("Temperature in Fº: ")
	fmt.Scanf("%f\n", &fahrenheit)

	celcius := 5 * ((fahrenheit - 32) / 9)

	fmt.Println(fmt.Sprintf("%vFº = %vCº", fahrenheit, celcius))
}
