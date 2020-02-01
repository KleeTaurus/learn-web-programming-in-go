package main

import (
	"fmt"
	"log"
	"net/http"
)

// This example will show how to create basic logging middleware in Go.
// A middleware simply takes a http.HanlerFunc as one of its parameters,
// wraps it and returns a new http.HanlerFunc for the server to call.

// The HandlerFunc type is an adapter to allow the use of ordinary functions
// as HTTP handlers. If f is a function with the appropriate signature,
// HandlerFunc(f) is a Handler that calls f.
// type HandlerFunc func(ResponseWriter, *Request)
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)

		// Calling the original handler
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {
	// We enhanced the foo and bar functions with the logging middleware.
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	// You can test the logging middleware by the following commands in your console.
	// curl -s http://localhost/foo
	// curl -s http://localhost/bar
	//
	// You will see the output as following in your console.
	// 2020/02/02 00:21:51 /foo
	// 2020/02/02 00:21:54 /bar
	http.ListenAndServe(":80", nil)
}
