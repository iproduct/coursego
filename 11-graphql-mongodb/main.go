package main

import (
	"github.com/go-chi/chi"
	"github.com/iproduct/coursego/11-graphql-mongodb/book"
	"github.com/iproduct/coursego/11-graphql-mongodb/infrastructure"
	"log"
	"net/http"
	"net/url"
)


func main() {
	routes := chi.NewRouter()
	r := book.RegisterRoutes(routes)
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
	env.InitMongoDB()
}
