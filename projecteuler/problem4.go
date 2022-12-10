package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(str string) bool {

	reversedStr := ""
	for index := len(str) - 1; index >= 0; index-- {
		reversedStr += string(str[index])
	}
	return str == reversedStr
}

func main() {

	var (
		base        int = 101
		limit       int = 1000
		palindromes []int
	)

	higher := 0
	for base < limit {
		factor := 101
		for factor < limit {
			product := base * factor
			if isPalindrome(strconv.Itoa(product)) {
				if product > higher {
					higher = product
				}
				palindromes = append(palindromes, product)
			}
			factor += 1
		}
		base += 1
	}

	fmt.Printf("Higher: %v\n", higher)

}
