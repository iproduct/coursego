package main

import (
	"encoding/json"
	"flag"
	"github.com/iproduct/coursego/fmi-2025-01-intro/books"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "server -addr :8080")

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	result, err := json.MarshalIndent(books.GoBooks, "", "    ")
	if err != nil {
	}
	w.Write(result)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/books", getBooks)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.
	log.Println("Starting HTTP server on", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
