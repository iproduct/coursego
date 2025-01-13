package main

import (
	gqhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/iproduct/coursego/11-graphql-simple-lab/schema"
	"log"
	"net/http"
)

func main() {
	handler := gqhandler.New(&gqhandler.Config{
		Schema:   &schema.Schema,
		GraphiQL: true,
		Pretty:   true,
	})

	http.Handle("/graphql", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
