package objects

import (
	"github.com/diegofu/iex-graphql/helpers"
	"github.com/diegofu/iex-graphql/types"
	"github.com/graphql-go/graphql"
)

var chartData = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "ChartData",
		Fields: helpers.Transform(&types.Chart{}),
	},
)

var Chart = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Chart",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: graphql.NewList(chartData),
			},
			"range": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
