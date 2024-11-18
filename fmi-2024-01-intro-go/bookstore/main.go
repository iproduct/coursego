package main

import (
	"flag"
	"github.com/iproduct/coursego/09-http/bookstore/books"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)

// ResourcesPath ia basic path to the project in filesystem
var workDir, _ = os.Getwd()

//var  ResourcesPath = path.Join(workDir, "bookstore")

var tmplAllBooks *template.Template

type database map[string]books.Book

type model struct {
	Db, Fav database
}

// var tmplAllBooks = template.Must(template.New("all-books").Parse(tmplAllBooksStr))
var db database = make(map[string]books.Book, 10) // books database is read only so no mutex is needed
var favourites database = make(map[string]books.Book, 10)
var rwlock sync.RWMutex                                         // Read-Write mutex defending access to favourites
var addr = flag.String("addr", ":8080", "http service address") // Q=17, R=18

func init() {
	tmplAllBooks = template.New("allBooks").Funcs(
		template.FuncMap{
			"urlSafe": func(url url.URL) template.HTML {
				return template.HTML(url.String())
			},
		})

	tmplAllBooks, err := tmplAllBooks.ParseFiles(
		path.Join(workDir, "templates", "head.html"),
		path.Join(workDir, "templates", "nav.html"),
		path.Join(workDir, "templates", "books.html"),
		path.Join(workDir, "templates", "favs.html"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(tmplAllBooks)
	log.Println(workDir)

	for _, book := range books.GoBooks {
		db[book.ID] = book
	}
}

func showBooks(w http.ResponseWriter, req *http.Request) {
	addFav := req.FormValue("add")
	removeFav := req.FormValue("remove")
	if addFav != "" {
		rwlock.Lock()
		favourites[addFav] = db[addFav]
		rwlock.Unlock()
		log.Printf("Book ID=%s aded to favourites\n", addFav)
	} else if removeFav != "" {
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
	rwlock.RLock()
	err := tmplAllBooks.ExecuteTemplate(w, "favs.html", favourites)
	rwlock.RUnlock()
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", showBooks)
	http.HandleFunc("/favs", showFavs)
	fs := http.FileServer(http.Dir(path.Join(workDir, "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Starting the HTTP server...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
