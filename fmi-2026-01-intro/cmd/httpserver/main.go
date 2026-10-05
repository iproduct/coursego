package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
func headers(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%v: %v <br>\n", name, value)
		}
	}
	scanner := bufio.NewScanner(req.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Fprintln(w, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(w, os.Stderr, "reading standard input:", err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("Starting server at port 8001")
	err := http.ListenAndServe(":8001", nil)
	log.Fatal(err)
}
