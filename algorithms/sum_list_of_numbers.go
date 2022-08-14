package main

import io "fmt"

func SumListOfNumbers(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + SumListOfNumbers(numbers[1:])
}

func main() {

	list := []int{1, 2, 3}

	io.Println(io.Sprintf("sum of %v = %v", list, SumListOfNumbers(list)))

}
