package main

type ID int
type Name string

func main() {

	var (
		name Name = "Jhon Doe"
		id   ID   = 1
	)

	println("Id: ", id)
	println("Name: ", name)
}
