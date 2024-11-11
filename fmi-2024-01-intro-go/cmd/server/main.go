package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Golang!")
}

func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for key, value := range r.Header {
		fmt.Fprintf(w, "%v: %v\n", key, value)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	log.Println("String HTTP Server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
