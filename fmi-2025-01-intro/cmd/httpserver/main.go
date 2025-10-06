package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%v: %v\n", name, value)
		}
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("server start")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
