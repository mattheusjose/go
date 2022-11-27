package main

import "fmt"

func main() {

	type Name string
	var (
		name Name = "Jhon Doe"
	)

	fmt.Printf("Type of variable name: %T\n", name)
	fmt.Printf("Value of variable name: %v", name)
}
