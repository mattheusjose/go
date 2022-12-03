package main

import "fmt"

type Address struct {
	city         string
	state        string
	neighborhood string
	number       int
	country      string
	zipCode      string
}

type Customer struct {
	name          string
	age           int
	taxIdentifier string
	phoneNumber   string
	address       Address
}

func main() {

	customer := Customer{
		name:          "Jhon Doe",
		age:           22,
		taxIdentifier: "76894809743",
		phoneNumber:   "5515996758332",
		address: Address{
			city:         "Itapeva",
			neighborhood: "Vila Amelia",
			number:       355,
			state:        "SP",
			zipCode:      "01001000",
			country:      "BRA",
		},
	}

	fmt.Println(customer)
}
