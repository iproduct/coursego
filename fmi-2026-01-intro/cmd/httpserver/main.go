package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
func headers(w http.ResponseWriter, req *http.Request) {
	// 1. Read request body
	text, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, fmt.Sprintln("Error reading request body:", err), http.StatusInternalServerError)
		return
	}

	// 2. echo headers and body
	w.Header().Add("Content-Type", "text/html")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%v: %v <br>\n", name, value)
		}
	}
	fmt.Fprintf(w, "%s\n", text)

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("Starting server at port 8001")
	err := http.ListenAndServe(":8001", nil)
	log.Fatal(err)
}
