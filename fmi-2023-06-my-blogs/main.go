package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/dao/inmemory"
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/domain/blogapp"
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/model"
	"html/template"
	"log"
	"net/http"
	"time"
)

type webapp struct {
	server         *http.Server
	mux            *http.ServeMux
	blog           *blogapp.BlogApp
	templateIndex  *template.Template
	templateCreate *template.Template
}

func (w *webapp) Run() error {
	w.mux.HandleFunc("/", w.handleMain)
	w.mux.HandleFunc("/create", w.handleCreate)
	w.mux.HandleFunc("/post", w.handlePost)
	w.mux.HandleFunc("/delete", w.handleDelete)
	w.templateIndex = template.Must(template.ParseFiles("./templates/index.tmpl.html"))
	w.templateCreate = template.Must(template.ParseFiles("./templates/create.tmpl.html"))
	log.Printf("server is listening at %s\n", w.server.Addr)
	if err := w.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start service on port %s:%w", w.server.Addr, err)
	}
	return nil
}

func (w *webapp) handleMain(writer http.ResponseWriter, request *http.Request) {
	posts, err := w.blog.GetAll()
	if err != nil {
		log.Printf("failed to get posts: %w", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	w.templateIndex.Execute(writer, posts)
}

func (w *webapp) handleCreate(writer http.ResponseWriter, request *http.Request) {
	w.templateCreate.Execute(writer, struct{}{})
}

func (w *webapp) handlePost(writer http.ResponseWriter, request *http.Request) {
	defer http.Redirect(writer, request, "/", http.StatusFound)
	err := request.ParseForm()
	if err != nil {
		log.Printf("error parsing html form: %w", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	post := model.Post{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Title:     request.Form.Get("title"),
		Content:   request.Form.Get("content"),
		Author:    request.Form.Get("author"),
	}

	if err = w.blog.Add(&post); err != nil {
		log.Printf("error parsing html form: %w", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (w *webapp) handleDelete(writer http.ResponseWriter, request *http.Request) {
	defer http.Redirect(writer, request, "/", http.StatusFound)

	idParams, ok := request.URL.Query()["id"]
	if !ok || len(idParams) == 0 {
		log.Printf("id is required")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id := idParams[0]
	err := w.blog.Delete(id)
	if err != nil {
		log.Printf("error parsing html form: %w", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
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
