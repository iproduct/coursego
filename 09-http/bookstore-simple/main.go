package main

import (
	"fmt"
	"log"
	"net/http"
)

var db database = make(map[string]Book, 10)

func init() {
	for _, book := range GoBooks {
		db[book.ID] = book
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type database map[string]Book

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for _, book := range db {
		fmt.Fprintf(w, "%s: %s - $%6.2f\n", book.ID, book.Title, book.RetailPrice)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if book, ok := db[id]; ok {
		fmt.Fprintf(w, "$%6.2f\n", book.RetailPrice)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no book with ID: %q\n", id)
	}
}

