package main

import (
	"math/rand"
	"net/http"
	"github.com/graphql-go/graphql"
	gqhandler "github.com/graphql-go/graphql-go-handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
		"postsCount": &graphql.Field {
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return rand.Intn(100), nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqhandler.New(&gqhandler.Config{
		Schema: &Schema,
		Pretty: true,
		GraphiQL: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
