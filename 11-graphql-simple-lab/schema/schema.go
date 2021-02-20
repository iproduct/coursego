package schema

import (
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello GraphQL!", nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
