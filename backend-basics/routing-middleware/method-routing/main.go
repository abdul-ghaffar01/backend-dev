package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from api")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Name string
		age  int
	}{"Abdul Ghaffar", 22}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Name string `json"name"`
		Age  int    `json:"age"`
	}{}
	json.NewDecoder(r.Body).Decode(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getUser(w, r)
		case "POST":
			createUser(w, r)
		}
	})
	http.ListenAndServe(":8000", nil)
}
