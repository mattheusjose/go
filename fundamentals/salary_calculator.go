package main

import (
	"fmt"
)

func CalculateSalary(amount float32, hours int) float32 {

	var salary float32 = amount * float32(hours)
	return salary
}

func main() {

	var amountPerHour float32
	var hoursWorked int

	fmt.Print("Amount per hour: R$")
	fmt.Scanf("%f\n", &amountPerHour)

	fmt.Print("Hours worked this month: ")
	fmt.Scanf("%d\n", &hoursWorked)

	salary := CalculateSalary(amountPerHour, hoursWorked)

	fmt.Println("Amount per hour: R$", amountPerHour)
	fmt.Println(fmt.Sprintf("Hours worked: %vh", hoursWorked))
	fmt.Println("Salary: R$", salary)
}
