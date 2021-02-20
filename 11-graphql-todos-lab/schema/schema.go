package schema

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/iproduct/coursego/11-graphql-todos-lab/model"
)

//GraphQL sSchema types and resolvers

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"todo": &graphql.Field{
			Type:        todoType,
			Description: "Get single todo.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: len(model.TodoList) - 1,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				idQuery, ok := params.Args["id"].(string)
				if ok {
					for _, todo := range model.TodoList {
						if todo.ID == idQuery {
							return todo, nil
						}
					}
				}
				return model.Todo{}, nil
			},
		},
		"lastTodo": &graphql.Field{
			Type:        todoType,
			Description: "Get last todo.",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return model.TodoList[len(model.TodoList)-1], nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "list of all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return model.TodoList, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        todoType,
			Description: "Create new todo.",
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				text, ok := params.Args["text"].(string)
				if ok {
					newTodo := model.Todo{
						ID:   uuid.New().String(),
						Text: text,
						Done: false,
					}
					model.TodoList = append(model.TodoList, newTodo)
					return newTodo, nil
				}
				return nil, fmt.Errorf("error updating todo: %v", params.Args["text"])
			},
		},
		"update": &graphql.Field{
			Type:        todoType,
			Description: "Update todo done status.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"done": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Boolean),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok1 := params.Args["id"].(string)
				done, ok2 := params.Args["done"].(bool)
				if ok1 && ok2 {
					// Search list for todo with id and change the done status
					for i := 0; i < len(model.TodoList); i++ {
						if model.TodoList[i].ID == id {
							model.TodoList[i].Done = done
							return model.TodoList[i], nil
						}
					}
					return nil, fmt.Errorf("error updating todo: id='%s' not found", id)
				}
				return nil, fmt.Errorf("error updating todo with id='%v'", params.Args["id"])
			},
		},
	},
})

var TodoSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
