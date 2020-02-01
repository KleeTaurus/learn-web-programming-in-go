package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User struct hold three fields: Firstname, Lastname and Age.
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		// NewDecoder returns a new decoder that reads from r.
		// Decode reads the next JSON-encoded value from its input and
		// stores it in the value pointed to by user.
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		// NewEncoder returns a new encoder that write to w.
		// Encode writes the JSON encoding of user (interface) to the stream,
		// followed by a newline character.
		// TODO We can pass the value or the pointer of user to Encode
		// method ??
		json.NewEncoder(w).Encode(user)
	})

	// You can test these two functions with following command in your console.

	// curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost/decode
	// Elon Musk is 48 years old!

	// curl -s http://localhost/encode
	// {"firstname":"John","lastname":"Doe","age":25}
	http.ListenAndServe(":80", nil)
}
