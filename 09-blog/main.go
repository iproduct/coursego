package main

import (
	"fmt"
	"github.com/iproduct/coursego/09-blog/blog"
	"github.com/iproduct/coursego/09-blog/container"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type rest struct {
	server *http.Server
	mux    *http.ServeMux
	blog   *blog.Blog
}

func (s *rest) Run() error {
	s.mux.HandleFunc("/", s.handleMain)
	s.mux.HandleFunc("/create", s.handleCreate)
	s.mux.HandleFunc("/post", s.createPost)
	s.mux.HandleFunc("/delete", s.deletePost)

	log.Printf("server is listening at %s\n", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start service on port %s: %w", s.server.Addr, err)
	}
	return nil
}

func (s *rest) handleMain(w http.ResponseWriter, r *http.Request) {
	posts, err := s.blog.GetAll()
	if err != nil {
		log.Panicln("failed to get posts: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	t := template.Must(template.ParseFiles("./templates/index.tmpl.html"))
	t.Execute(w, posts)
}

func (s *rest) handleCreate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/create.tmpl.html"))
	t.Execute(w, struct{}{})
}

func (s *rest) createPost(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	r.ParseForm()
	post := blog.Post{
		ID:        uuid.New().String(),
		Content:   r.Form.Get("content"),
		Author:    r.Form.Get("author"),
		Heading:   r.Form.Get("heading"),
		CreatedAt: time.Now(),
	}

	err := s.blog.NewPost(&post)
	if err != nil {
		log.Printf("failed to insert post: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *rest) deletePost(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	idParams, ok := r.URL.Query()["id"]
	if !ok || len(idParams) == 0 {
		log.Printf("id is required")
		return
	}
	id := idParams[0]
	err := s.blog.DeletePost(id)
	if err != nil {
		log.Printf("failed to delete post: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// container := container.NewMySQLStore(container.MySQLOptions{
	// 	URI: fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/09-blog?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS")),
	// })
	container := container.NewMongoStore(container.MongoOptions{
		URI: "mongodb://localhost:27017",
	})
	err := container.Init()
	if err != nil {
		log.Fatalf("failed to init store: %s", err)
	}
	// container := container.NewInMemory()
	blog := blog.New(&container)

	rest := rest{
		server: server,
		mux:    mux,
		blog:   blog,
	}
	rest.Run()
}
