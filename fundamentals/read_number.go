package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Print("Type a number: ")
	fmt.Scanf("%d", &number)

	fmt.Println(fmt.Sprintf("The number is %v", number))
}
