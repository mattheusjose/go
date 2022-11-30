package main

import (
	"errors"
	"fmt"
)

func main() {

	functions := map[string]func(int, int) (int, error){
		"add": func(n1 int, n2 int) (int, error) {

			return n1 + n2, nil
		},
		"sub": func(a int, b int) (int, error) {
			return a - b, nil
		},
		"mul": func(n1 int, n2 int) (int, error) {

			return n1 * n2, nil
		},
		"div": func(divisor int, div int) (int, error) {

			if div == 0 {
				return 0, errors.New("divident must be greather than zero")
			}

			return divisor / div, nil
		},
	}

	fmt.Println("[add] - To sum")
	fmt.Println("[sub] - To subtract")
	fmt.Println("[mul] - To multiply")
	fmt.Println("[div] - To divide: ")
	var option string = "div"

	action := functions[option]

	n1 := 10
	n2 := 2

	res, error := action(n1, n2)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println("Res: ", res)
}
