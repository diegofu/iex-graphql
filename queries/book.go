package queries

import (
	"fmt"

	"github.com/diegofu/iex-graphql/requests"

	"github.com/diegofu/iex-graphql/objects"
	"github.com/graphql-go/graphql"
)

var Book = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "BookQuery",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type: objects.Book,
				Args: graphql.FieldConfigArgument{
					"symbol": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					symbol, ok := p.Args["symbol"].(string)

					if !ok {
						return nil, fmt.Errorf("Unable to parse args %v", p.Args["symbol"])
					}

					book, err := requests.GetBook(symbol)
					if err != nil {
						return nil, err
					}

					return book, nil
				},
			},
		},
	},
)
