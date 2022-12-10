package main

import "fmt"

func main() {

	var (
		limit          int = 1000
		threeMultiple  int = 3
		fiveMultiple   int = 5
		sumOfMultiples int = 0
	)

	for index := 1; index < limit; index++ {

		if index%threeMultiple == 0 || index%fiveMultiple == 0 {
			sumOfMultiples += index
		}
	}

	fmt.Printf("The sum of all mutiples of %v and %v bellow %v is %v",
		threeMultiple,
		fiveMultiple,
		limit,
		sumOfMultiples,
	)
}
