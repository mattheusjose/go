package main

import (
	"fmt"
)

func NumInList(list []int, target int) bool {

	for _, number := range list {
		if number == target {
			return true
		}
	}

	return false
}

func main() {

	list := []int{1, 2, 3, 4, 5, 6}
	targetA := 4
	targetB := 10

	fmt.Println(fmt.Sprintf("%v in %v = %v", targetA, list, NumInList(list, targetA)))
	fmt.Println(fmt.Sprintf("%v in %v = %v", targetB, list, NumInList(list, targetB)))
}
