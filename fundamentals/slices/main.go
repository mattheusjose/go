package main

import "fmt"

func main() {

	array := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(array), cap(array), array)
	fmt.Printf("len=%d cap=%d %v\n", len(array[:0]), cap(array[:0]), array[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(array[:4]), cap(array[:4]), array[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(array[2:]), cap(array[2:]), array[2:])
}
