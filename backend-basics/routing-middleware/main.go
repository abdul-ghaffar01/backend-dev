package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from api")
}

func main(){
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}