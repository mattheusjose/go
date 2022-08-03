package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var users = map[string]map[string]string{}

type ResultData struct {
	users map[string]map[string]string
}

func loadTemplate(response http.ResponseWriter, request *http.Request) {

	template, err := template.ParseFiles("./static/users.html")

	if err != nil {
		fmt.Fprintf(response, "Error parsing json")
		return
	}

	err = template.Execute(response, users)
}

func createUsers(response http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
		return
	}

	name := request.FormValue("name")
	email := request.FormValue("email")
	key := strings.ToLower(strings.Join(strings.Split(name, " "), ""))

	users[key] = make(map[string]string)
	users[key]["id"] = "2"
	users[key]["name"] = name
	users[key]["email"] = email

	loadTemplate(response, request)
	http.Redirect(response, request, "/users.html", http.StatusSeeOther)
}

func main() {

	const port = ":8080"
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/api/v1/users/create", createUsers)

	log.Printf("Starting server at port %s", port)
	serverConnection := http.ListenAndServe(port, nil)

	if serverConnection != nil {
		log.Fatal(serverConnection)
	}

}
