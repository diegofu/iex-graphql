package objects

import (
	"github.com/graphql-go/graphql"
)

var Book = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"quote": &graphql.Field{
				Type: Quote,
			},
		},
	},
)
