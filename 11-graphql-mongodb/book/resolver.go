package book

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"bookById": &graphql.Field{
				Type:        productType,
				Description: "Get book by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					id, ok := p.Args["id"].(string)
					if ok {
						// Find product
						result = GetBookByID(context.Background(), id)
					}
					return result, nil
				},
			},
			"bookByName": &graphql.Field{
				Type:        productType,
				Description: "Get book by name",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					name, ok := p.Args["name"].(string)
					if ok {
						// Find product
						result = GetBookByName(context.Background(), name)
					}
					return result, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get book list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 10,
						Description:  "Number of books to fetch",
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					limit, _ := params.Args["limit"].(int)
					result = GetBookList(context.Background(), limit)
					return result, nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        productType,
			Description: "Create new book",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				book := Book{
					Name:        params.Args["name"].(string),
					Description: params.Args["description"].(string),
					Price:       params.Args["price"].(float64),
				}
				if err := InsertBook(context.Background(), &book); err != nil {
					return nil, err
				}

				return book, nil
			},
		},

		"update": &graphql.Field{
			Type:        productType,
			Description: "Update book by name",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, idOk := params.Args["id"].(string)
				if !idOk {
					return nil, fmt.Errorf("error updating book: ID is missing")
				}
				book := Book{
					ID: id,
				}
				if name, nameOk := params.Args["name"].(string); nameOk {
					book.Name = name
				}
				if price, priceOk := params.Args["price"].(float64); priceOk {
					book.Price = price
				}
				if description, descriptionOk := params.Args["description"].(string); descriptionOk {
					book.Description = description
				}

				if err := UpdateBook(context.Background(), book); err != nil {
					return nil, err
				}
				return book, nil
			},
		},

		"delete": &graphql.Field{
			Type:        productType,
			Description: "Delete book by name",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["id"].(string)
				return DeleteBook(context.Background(), name)
			},
		},
	},
})

// schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
