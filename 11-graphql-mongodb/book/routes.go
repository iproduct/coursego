package book

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-go/handler"
)

func RegisterRoutes(r *chi.Mux) *chi.Mux {
	/* GraphQL */
	graphQL := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Use(middleware.Logger)
	r.Handle("/graphql", graphQL)

	/* Rest API */

	r.Get("/books", RestApiGetBookAllBooks)
	r.Get("/books/{bookname}", RestApiGetBook)
	return r
}
