package main

import (
	"fmt"
)

func main() {

	var a interface{} = false
	var b interface{} = 10
	var c interface{} = "Hello, World!"

	showType(a)
	showType(b)
	showType(c)
}

func showType(x interface{}) {

	fmt.Printf("O tipo de x é %T e o valor é %v\n", x, x)
}
