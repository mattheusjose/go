package main

import "fmt"

type Number int

func printArray(array []Number) {

	for _, element := range array {
		fmt.Printf("[%v] ", element)
	}
}

func main() {

	array := []Number{24, 57, 864, 874, 357, 386}
	printArray(array)
}
