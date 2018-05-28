package objects

import (
	"github.com/graphql-go/graphql"
)

var Quote = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Quote",
		Fields: graphql.Fields{
			"symbol": &graphql.Field{
				Type: graphql.String,
			},
			"companyName": &graphql.Field{
				Type: graphql.String,
			},
			"primaryExchange": &graphql.Field{
				Type: graphql.String,
			},
			"open": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
