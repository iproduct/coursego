package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!!!")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, value := range r.Header {
		for _, h := range value {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("Server is running ...")
	http.ListenAndServe(":8080", nil)
}
