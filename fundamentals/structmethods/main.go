package main

import (
	"fmt"
)

type Address struct {
	city         string
	state        string
	neighborhood string
}

type Customer struct {
	name    string
	age     int
	enabled bool
	address Address
}

func (c *Customer) disableCustomer() {
	fmt.Println("Disable customer...")
	c.enabled = false
}

func (c *Customer) enableCustomer() {
	fmt.Println("Enable customer...")
	c.enabled = true
}

func main() {

	customer := Customer{
		name:    "Jhon Doe",
		enabled: true,
		age:     22,
	}

	customer.disableCustomer()
	customer.enableCustomer()
	fmt.Println(customer)
}
