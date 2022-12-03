package main

import (
	"fmt"
)

type Customer struct {
	name    string
	age     int
	enabled bool
}

func printCustomer(customer Customer) {

	fmt.Printf("Name: %v\n", customer.name)
	fmt.Printf("Age: %v\n", customer.age)
	fmt.Printf("Is enabled: %v\n", customer.enabled)
}

func increaseAge(customer *Customer) {
	customer.age = customer.age + 1
}

func main() {

	c1 := Customer{
		name:    "Jhon Doe",
		age:     22,
		enabled: false,
	}

	increaseAge(&c1)
	printCustomer(c1)
}
