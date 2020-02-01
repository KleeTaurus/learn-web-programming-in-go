package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// In this example you will see how to use the gorilla/mux package to create routes with
// named parameters, GET/POST handlers and domain restrictions.

// gorilla/mux is a package which adapts to Go's default HTTP router, use the go get command
// to install the package from GitHub like so:
// go get -u github.com/gorilla/mux

func main() {
	// First create a new request router. The router is the main router for your web
	// application and will later be passed as parameter to the server. It will receive
	// all HTTP connections and pass it on to the request handlers you will register on it.
	// You can create a new router like so:
	r := mux.NewRouter()

	// Once you have a new router you can register request handlers like usual.
	// To have a request handler match the URL you replace the dynamic segments of with
	// placeholders in your URL pattern like so:
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		// The last thing is to get the data from these segments. The package comes with
		// the function mux.Vars(r) which takes the http.Request as parameter and returns
		// a map of the segments.
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"]   // the page

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// The second parameter of http.ListenAndServe means to use the default router of the
	// net/http package. To make use of your own router, replace the nil with the variable
	// of your router r.

	// You can navigate your browser to http://localhost/books/go-programming-blueprint/page/10
	// and see your server handling your request.
	http.ListenAndServe(":80", r)
}

// Other features of the gorilla/mux Router

// 1. Restrict the request handler to specific HTTP methods.
// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

// 2. Restrict the request handler to specific hostnames or subdomains
// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

// 3. Restrict the request handler to specific http/https.
// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

// 4. Restrict the reqeust handler to specific path prefixes.
// bookrouter := r.PathPrefix("/books").Subrouter()
// bookrouter.HandleFunc("/", AllBooks)
// bookrouter.HandleFunc("/{title}", GetBook)
