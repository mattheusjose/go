package main

import (
	"fmt"
)

func main() {

	var totalFishedWeight int
	const tax = 4
	const limitToPayTax = 50
	fmt.Println("Type total of fish (KG): ")
	fmt.Scanf("%d", &totalFishedWeight)

	totalExceed := totalFishedWeight - limitToPayTax
	if totalExceed < 0 {
		totalExceed = 0
	}

	fee := totalExceed * tax

	fmt.Println("Total fished (KG): ", totalFishedWeight)
	fmt.Println("Total exceed (KG): ", totalExceed)
	fmt.Println("Fee: R$", fee)
}
