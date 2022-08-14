package main

import (
	io "fmt"
)

func FizzBuzz(number int) {

	for index := 1; index <= number; index++ {

		text := io.Sprintf("%v", index)

		if index%15 == 0 {
			text = "Fizz Buzz"
		} else if index%3 == 0 {
			text = "Fizz"
		} else if index%5 == 0 {
			text = "Buzz"
		}

		io.Print(text)

		if index < number {
			io.Print(", ")
		}
	}
}

func main() {

	io.Println("-=-=-=-=-=-=-=-=-=-=-=-= Fizz Buzz -=-=-=-=-=-=-=-=-=-=")

	FizzBuzz(20)
}
