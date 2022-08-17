package main

import (
	"fmt"
)

func Sum(n1 int, n2 int) int {

	res := n1 + n2
	return res
}

func main() {

	var n1 int
	var n2 int

	fmt.Print("Type the first number: ")
	fmt.Scanf("%d\n", &n1)

	fmt.Print("Type the second number: ")
	fmt.Scanf("%d", &n2)

	res := Sum(n1, n2)
	fmt.Println(fmt.Sprintf("%v + %v = %v", n1, n2, res))
}
