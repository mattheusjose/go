package main

import "fmt"

func main() {

	var (
		n1 int = 10
		n2 int = 20
		n3 int = 100
	)

	res := sum(n1, n2, n3)
	fmt.Println("Res: ", res)
}

func sum(numbers ...int) int {

	res := 0
	for _, number := range numbers {
		res += number
	}

	return res
}
