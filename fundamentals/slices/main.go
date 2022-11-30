package main

import "fmt"

func main() {

	array := []int{10, 20, 30, 40, 50, 60, 70, 80, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(array), cap(array), array)
	fmt.Printf("len=%d cap=%d %v\n", len(array[:0]), cap(array[:0]), array[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(array[:4]), cap(array[:4]), array[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(array[2:]), cap(array[2:]), array[2:])
	array = append(array, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(array), cap(array), array)
	fmt.Printf("len=%d cap=%d %v\n", len(array[2:]), cap(array[2:]), array[2:])
}
