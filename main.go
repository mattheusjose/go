package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(response http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(response, "POST request sucessfull")
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(response, "Name %s\n", name)
	fmt.Fprintf(response, "Address %s\n", address)
}

func handleHello(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "Not found", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(response, "Method is not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(response, "Hello World Go!")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running at port 8080")
}
