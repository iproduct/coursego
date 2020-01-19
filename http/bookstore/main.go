package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"
)

const ResourcesPath = "D:/CourseGO/workspace/src/github.com/iproduct/coursego/http/bookstore"

var tmplBase = template.New("base").Funcs(
	template.FuncMap{
		"urlSafe": func(url url.URL) template.HTML {
			return template.HTML(url.String())
		},
	})

var tmplAllBooks, _ = tmplBase.ParseFiles(path.Join(ResourcesPath, "templates", "books.html"))
//var tmplAllBooks = template.Must(template.New("all-books").Parse(tmplAllBooksStr))
var db database = make(map[string]Book, 10)
var favourites database = make(map[string]Book, 10)
var addr = flag.String("addr", ":8080", "http service address") // Q=17, R=18

func init() {
	for _, book := range goBooks {
		db[book.ID] = book
	}
}
type database map[string]Book

type model struct{
	Db,	Fav database
}

func books(w http.ResponseWriter, req *http.Request) {
	addFav := req.FormValue("add")
	favourites[addFav] = db[addFav]
	log.Printf("Book ID=%s aded to favourites\n", addFav)
	data := model{Db: db, Fav: favourites}
	err := tmplAllBooks.ExecuteTemplate(w, "books.html", data)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
	}
	//for _, book := range db {
	//	fmt.Fprintf(w, "%s: %s - $%6.2f\n", book.ID, book.Title, book.RetailPrice)
	//}
}

func price(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if book, ok := db[id]; ok {
		fmt.Fprintf(w, "$%6.2f\n", book.RetailPrice)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no book with ID: %q\n", id)
	}
}

func main() {
	http.HandleFunc("/books", books)
	http.HandleFunc("/price", price)
	fs := http.FileServer(http.Dir(path.Join(ResourcesPath, "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Starting the HTTP server ...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}