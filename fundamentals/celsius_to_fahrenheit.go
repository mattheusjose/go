package main

import (
	"fmt"
)

func main() {

	var celsius float32
	fmt.Print("Type temperature in Cº: ")
	fmt.Scanf("%f\n", &celsius)

	fahrenheit := celsius*9/5 + 32

	fmt.Println(fmt.Sprintf("%vCº = %vFº", celsius, fahrenheit))

}
