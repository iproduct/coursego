// Writing a basic HTTP simple-server is easy using the
// `net/09-http` package.
package main

import (
	"encoding/json"
	"flag"
	"github.com/iproduct/coursego/09-http/server-json/books"
	"log"
	"net/http"
)
var addr = flag.String("addr", ":8080", "server -addr :8080")

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(books.GoBooks)
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
