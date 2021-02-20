package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	gqhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/iproduct/coursego/11-graphql-todos-lab/model"
	"github.com/iproduct/coursego/11-graphql-todos-lab/schema"
	"log"
	"net/http"
)

func init() {
	todo1 := model.Todo{ID: "a", Text: "A todo not to forget", Done: false}
	todo2 := model.Todo{ID: "b", Text: "This is the most important", Done: false}
	todo3 := model.Todo{ID: "c", Text: "Please do this or else", Done: false}
	model.TodoList = append(model.TodoList, todo1, todo2, todo3)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("error executing graphql query '%v' : %v\n", result.Errors)
	}
	return result
}

func main() {
	http.Handle("/graphql", gqhandler.New(&gqhandler.Config{
		Schema: &schema.TodoSchema,
		Pretty: true,
	}))
	//http.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
	//	result := executeQuery(request.URL.Query().Get("query"), schema.TodoSchema)
	//	json.NewEncoder(writer).Encode(result)
	//})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", fs)
	fmt.Printf("Server is running on port 8080\n")
	fmt.Printf("TODOS: %v", executeQuery("{list{id text done}}", schema.TodoSchema))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
