package main

import (
	"net/http"
	"log"
	"fmt"
)

type identity struct {
	firstName string
	lastName string
}

const (
	pageTop = `
		<!DOCTYPE HTML>
		<html>
			<head>
				<title>Call Me!</title>
				<style>
					.error {
						color: #FF0000;
					}
				</style>
			</head>
			<body>
				<h1>Call Me!</h1>
				<p>How do you want to be called?</p>
	`
	pageForm = `
		<form action="/" method="POST">
			<label for="firstName">First name:</label>
			<br />
			<input type="text" name="firstName" size="30" />
			<br />
			<label for="lastName">Last name:</label>
			<br />
			<input type="text" name="lastName" size="30" />
			<br />
			<input type="submit" value="Call Me!">
		</form>
	`
	pageBottom = `</body></html>`
	pageError = `<p class="error">%s</p>`
)

func main() {
	run()
}

func run() {
	http.HandleFunc("/", renderPage)
	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func renderPage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, pageForm)
	if err != nil {
		fmt.Fprintf(writer, pageError, err)
	} else {
		if firstName, lastName, ok := processRequest(request); ok {
			id := setIdentity(firstName, lastName)
			res := callMe(id)
			fmt.Fprintf(writer, res)
		}
	}
	fmt.Fprintf(writer, pageBottom)
}

func processRequest(request *http.Request) (string, string, bool) {
	var (
		firstName string
		lastName string
	)

	if firstNameSlice, found := request.Form["firstName"]; found && len(firstNameSlice) > 0 {
		firstName = firstNameSlice[0]
	}
	
	if lastNameSlice, found := request.Form["lastName"]; found && len(lastNameSlice) > 0 {
		lastName = lastNameSlice[0]
	}

	if len(firstName) == 0 || len(lastName) == 0 {
		return "", "", false
	}
	
	return firstName, lastName, true
}

func setIdentity(firstName string, lastName string) (id identity) {
	id.firstName = firstName
	id.lastName = lastName
	return id
}

func callMe(id identity) string {
	return fmt.Sprintf(
		`<div>Hi %s %s!</div>`,
		id.firstName,
		id.lastName)
}
