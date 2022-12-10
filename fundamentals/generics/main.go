package main

import (
	"fmt"
)

type Integer int
type Decimal float32

type Number interface {
	~int | ~float32 | float64
}

func sum[T Number](a, b T) T {
	return a + b
}

func compare[T comparable](a, b T) bool {
	return a == b
}

func main() {

	var (
		a int     = 10
		b int     = 20
		c float32 = 20.5
		d float32 = 20.4
		e float64 = 19.5
		f float64 = 20.4
		g Integer = 10
		h Integer = 20
		i Decimal = 100.45
		j Decimal = 455.53
	)

	sum1 := sum(a, b)
	sum2 := sum(c, d)
	sum3 := sum(e, f)
	sum4 := sum(g, h)
	sum5 := sum(i, j)

	fmt.Println("Sum 1: ", sum1)
	fmt.Println("Sum 2: ", sum2)
	fmt.Println("Sum 3: ", sum3)
	fmt.Println("Sum 4: ", sum4)
	fmt.Println("Sum 5: ", sum5)

	fmt.Printf("%v == %v = %v\n", a, b, compare(a, b))
	fmt.Printf("%v == %v = %v\n", c, d, compare(c, d))
	fmt.Printf("%v == %v = %v\n", e, f, compare(e, f))
	fmt.Printf("%v == %v = %v\n", g, h, compare(g, h))
	fmt.Printf("%v == %v = %v\n", i, j, compare(i, j))
}
