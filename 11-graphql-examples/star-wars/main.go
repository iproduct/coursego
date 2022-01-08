package main

import (
	"net/http"

	gqhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/graphql-go/graphql/testutil"
)

func main() {
	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqhandler.New(&gqhandler.Config{
		Schema:   &testutil.StarWarsSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)

	//http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
	//	query := r.URL.Query().Get("query")
	//	result := graphql.Do(graphql.Params{
	//		Schema:        testutil.StarWarsSchema,
	//		RequestString: query,
	//	})
	//	json.NewEncoder(w).Encode(result)
	//})
	//fmt.Println("Now server is running on port 8080")
	//fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={hero{name}}'")
	//http.ListenAndServe(":8080", nil)
}
