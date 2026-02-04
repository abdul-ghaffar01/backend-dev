package main

import (
	"fmt"
	"net/http"
)

func helloHanlder (resp http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello world")
	}
func main() {
	http.HandleFunc("/", helloHanlder)
	http.ListenAndServe(":8000",nil)	
}