package main

import (
	"fmt"
	"net/http"
)

// In this example you will learn how to create a basic HTTP server in Go.

func main() {
	// We can register a new handler with the http.HandleFunc function.
	// It's first parameter takes a path to match and a function to
	// execute as a second.
	// For the dynamic aspect, the http.Request contains all information
	// about the request and it's parameters. You can read GET parameters
	// with r.URL.Query().Get("token") or POST parameters (fields from an
	// HTML form) with r.FormValue("email")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	// FileServer returns a handler that serves HTTP requests with the
	// contents of the file system rooted at root.
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// You can navigate your browser to http://localhost/static/example.js
	// and see your server handling your request.
	http.ListenAndServe(":80", nil)
}
