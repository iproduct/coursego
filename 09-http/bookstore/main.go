package main

import (
	"flag"
	"github.com/iproduct/coursego/09-http/bookstore/books"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"
)

// ResourcesPath ia basic path to the project in filesystem
const ResourcesPath = "D:/CourseGO/git/coursego/09-http/bookstore"

var tmplBase = template.New("base").Funcs(
	template.FuncMap{
		"urlSafe": func(url url.URL) template.HTML {
			return template.HTML(url.String())
		},
	})

var tmplAllBooks, _ = tmplBase.ParseFiles(
	path.Join(ResourcesPath, "templates", "books.html"),
	path.Join(ResourcesPath, "templates", "favs.html"),
)

//var tmplAllBooks = template.Must(template.New("all-books").Parse(tmplAllBooksStr))
var db database = make(map[string]books.Book, 10) // books database is read only so no mutex is needed
var favourites database = make(map[string]books.Book, 10)
var rwlock sync.RWMutex // Read-Write mutex defending access to favourites
var addr = flag.String("addr", ":8080", "http service address") // Q=17, R=18

func init() {
	for _, t := range tmplAllBooks.Templates() {
		t.ParseFiles(
			path.Join(ResourcesPath, "templates", "head.html"),
			path.Join(ResourcesPath, "templates", "nav.html"),
		)
	}
	for _, book := range books.GoBooks {
		db[book.ID] = book
	}
}

type database map[string]books.Book

type model struct {
	Db, Fav database
}

func showBooks(w http.ResponseWriter, req *http.Request) {
	addFav := req.FormValue("add")
	removeFav := req.FormValue("remove")
	if addFav != "" {
		rwlock.Lock()
		favourites[addFav] = db[addFav]
		rwlock.Unlock()
		log.Printf("Book ID=%s aded to favourites\n", addFav)
	}else if removeFav != "" {
		rwlock.Lock()
		delete(favourites, removeFav)
		rwlock.Unlock()
		log.Printf("Book ID=%s remove from favourites\n", removeFav)
	}

	data := model{Db: db, Fav: favourites}
	rwlock.RLock()
	err := tmplAllBooks.ExecuteTemplate(w, "books.html", data)
	rwlock.RUnlock()
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
	}
}

func showFavs(w http.ResponseWriter, req *http.Request) {
	removeFav := req.FormValue("remove")
	if removeFav != "" {
		rwlock.Lock()
		delete(favourites, removeFav)
		rwlock.Unlock()
		log.Printf("Book ID=%s remove from favourites\n", removeFav)
	}
	err := tmplAllBooks.ExecuteTemplate(w, "favs.html", favourites)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
	}
}



func main() {
	http.HandleFunc("/", showBooks)
	http.HandleFunc("/favs", showFavs)
	fs := http.FileServer(http.Dir(path.Join(ResourcesPath, "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Starting the HTTP server ...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
