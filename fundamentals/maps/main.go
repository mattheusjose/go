package main

import "fmt"

func main() {

	persons := map[string]map[string]string{
		"p1": {
			"name":   "Jhon Doe",
			"age":    "20",
			"single": "true",
		},
		"p2": {
			"name":   "Sara Doe",
			"age":    "29",
			"single": "false",
		},
	}

	for person := range persons {

		fmt.Println("-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		fmt.Printf("Name: %v\n", persons[person]["name"])
		fmt.Printf("Age: %v\n", persons[person]["age"])
		fmt.Printf("Single: %v\n", persons[person]["single"])
		fmt.Println("-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	}
}
