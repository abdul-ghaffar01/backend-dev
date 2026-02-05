package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("The method of request is %q\n", r.Method)
		startTime := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("Request completed in %d microseconds\n", (time.Now().Sub(startTime)).Microseconds())
	})
}

func HomeHanlder(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Handling the home route...")
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware) // will log each request
	r.HandleFunc("/", HomeHanlder)
	http.ListenAndServe(":8000", r)
}
