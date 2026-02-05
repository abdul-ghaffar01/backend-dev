package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{}

func getUserById(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	for _, user := range users {
		if user.Id == id {
			resp.Header().Set("Content-Type", "application/json")
			json.NewEncoder(resp).Encode(user)
			return
		}
	}
}

func getAllUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(users)

}

func createUser(resp http.ResponseWriter, req *http.Request) {
	var u User
	json.NewDecoder(req.Body).Decode(&u)
	u.Id = strconv.Itoa(len(users) + 1)
	users = append(users, u)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(u)
}

func main() {
	// To get user with id
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getUserById(w, r)
		case "POST":
			createUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getAllUsers(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8000", nil)
}
