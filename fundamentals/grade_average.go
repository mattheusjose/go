package main

import (
	"fmt"
)

func main() {

	var n1, n2, n3, n4 float32

	fmt.Print("Type the first grade note: ")
	fmt.Scanf("%f\n", &n1)

	fmt.Print("Type the second grade note: ")
	fmt.Scanf("%f\n", &n2)

	fmt.Print("Type the third grade note: ")
	fmt.Scanf("%f\n", &n3)

	fmt.Print("Type the fourth grade note: ")
	fmt.Scanf("%f\n", &n4)

	average := (n1 + n2 + n3 + n4) / 4

	fmt.Println("Your average is: ", average)
}
