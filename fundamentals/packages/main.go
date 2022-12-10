package main

import (
	"arithmetic/functions"
	"fmt"
)

func main() {

	fmt.Println("Arithmetic Operations")

	var (
		a int = 50
		b int = 20
	)

	sumRes := functions.Sum(a, b)
	subRes := functions.Subtract(a, b)
	divRes := functions.Divide(a, b)
	mulRes := functions.Multiply(a, b)

	fmt.Printf("Operation Id: %v", functions.RandomUUID())
	fmt.Printf("%v+%v=%v\n", a, b, sumRes)
	fmt.Printf("%v-%v=%v\n", a, b, subRes)
	fmt.Printf("%v/%v=%v\n", a, b, divRes)
	fmt.Printf("%vx%v=%v\n", a, b, mulRes)
}
