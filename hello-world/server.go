package main

import (
	"fmt"
	"net/http"
)

// In this example you will figure out how simple it is, to create
// a webserver that you can view in your browser.

func main() {
	// Registering a request handler to the default HTTP server is
	// as simple as this.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// The following code will start Go's default HTTP server and
	// listen for connections on port 80. You can navigate your
	// browser to http://localhost/ and see your server handling
	// your request.
	http.ListenAndServe(":80", nil)
}
