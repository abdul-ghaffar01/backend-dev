package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the api")
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	json.NewEncoder(w).Encode(struct {
		Id string		`json:"id"`
		Name string		`json:"name"`
		Age  int		`json:"age"`
	}{id, "Abdul ghaffar", 22})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/user/{id}", getUserById).Methods("GET", "POST")
	http.ListenAndServe(":8000", r)
}
