package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	zipCode := "18411515"
	fmt.Println(GetAddress(zipCode))
}

func GetAddress(zipCode string) string {

	url := "https://viacep.com.br/ws/" + zipCode + "/json/"
	request, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	response, error := io.ReadAll(request.Body)

	if error != nil {
		panic(error)
	}

	request.Body.Close()

	return string(response)
}
