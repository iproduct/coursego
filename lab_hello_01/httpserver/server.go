package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Golang!")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, value := range r.Header {
		for _, hval := range value {
			fmt.Fprintf(w, "%v: %v\n", name, hval)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("Staring HTTP server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
