package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// This example will show how to simulate a contact form and parse the message into a struct.

// ContactDetails will store the value submitted from user browser
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	// ParseFiles parses the forms.html file
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// Extracted the form's value and store them in a struct variable
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// Print details value on the console
		fmt.Printf("User submited: %+v\n", details)

		// Pass the anonymous struct value to template for html rendering
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":80", nil)
}
