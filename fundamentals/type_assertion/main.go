package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = 20

	var_a, isAInteger := x.(int)
	var_b, isBInteger := y.(int)

	if isAInteger {
		fmt.Println(var_a)
	}

	if isBInteger {
		fmt.Println(var_b)
	}
}

func sum(x interface{}, y interface{}) int {
	return x.(int) + y.(int)
}
