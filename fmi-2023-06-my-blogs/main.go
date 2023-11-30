package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/dao/inmemory"
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/domain/blogapp"
	"log"
	"net/http"
)

type webapp struct {
	server *http.Server
	mux    *http.ServeMux
	blog   *blogapp.BlogApp
}

func (w *webapp) Run() error {
	w.mux.HandleFunc("/", w.handleMain)
	w.mux.HandleFunc("/create", w.handleCreate)
	w.mux.HandleFunc("/post", w.handlePost)
	w.mux.HandleFunc("/delete", w.handleDelete)
	log.Printf("server is listening at %s\n", w.server.Addr)
	if err := w.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start service on port %s:%w", w.server.Addr, err)
	}
	return nil
}

func (w *webapp) handleMain(writer http.ResponseWriter, request *http.Request) {

}

func (w *webapp) handleCreate(writer http.ResponseWriter, request *http.Request) {

}

func (w *webapp) handlePost(writer http.ResponseWriter, request *http.Request) {

}

func (w *webapp) handleDelete(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	blogRepo := inmemory.New()
	blog := blogapp.New(blogRepo)
	app := webapp{
		server: server,
		mux:    mux,
		blog:   blog,
	}
	app.Run()
}
