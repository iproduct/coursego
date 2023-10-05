package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello from Golang")
	if err != nil {
		return
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, value := range r.Header {
		for _, hval := range value {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, hval)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("Starting HTTP server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
