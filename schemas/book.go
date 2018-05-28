package schemas

import (
	"github.com/diegofu/iex-graphql/queries"
	"github.com/graphql-go/graphql"
)

func NewBook() (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queries.Book,
		},
	)
}
