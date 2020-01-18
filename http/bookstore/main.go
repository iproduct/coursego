package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var db database = make(map[string]Book, 10)
var tmplAllBooks = template.Must(template.New("all-books").Parse(tmplAllBooksStr))

func init() {
	for _, book := range goBooks {
		db[book.ID] = book
	}
}

func main() {
	http.HandleFunc("/books", db.books)
	http.HandleFunc("/price", db.price)
	log.Println("Starting the HTTP server ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type database map[string]Book

func (db database) books(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	err := tmplAllBooks.Execute(w, goBooks)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
	}
	//for _, book := range db {
	//	fmt.Fprintf(w, "%s: %s - $%6.2f\n", book.ID, book.Title, book.RetailPrice)
	//}
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

