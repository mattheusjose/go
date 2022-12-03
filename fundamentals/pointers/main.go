package main

import "fmt"

func sum(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {

	var (
		a    int  = 10
		b    int  = 20
		ptrA *int = &a
		ptrB *int = &b
	)

	res := sum(ptrA, ptrB)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(res)

}
