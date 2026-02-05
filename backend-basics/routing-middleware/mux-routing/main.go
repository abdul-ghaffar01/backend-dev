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
		Id   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{id, "Abdul ghaffar", 22})
}

func infoOfV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the version 1 of the apis")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/user/{id}", getUserById).Methods("GET", "POST")

	// Sub routers
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", infoOfV1)
	http.ListenAndServe(":8000", r)
}
