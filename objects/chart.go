package objects

import (
	"github.com/graphql-go/graphql"
)

var chartData = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ChartData",
		Fields: graphql.Fields{
			"minute": &graphql.Field{
				Type: graphql.String,
			},
			"marketAverage": &graphql.Field{
				Type: graphql.Float,
			},
			"marketNational": &graphql.Field{
				Type: graphql.Float,
			},
			"marketNumberOfTrades": &graphql.Field{
				Type: graphql.Int,
			},
			"marketOpen": &graphql.Field{
				Type: graphql.Float,
			},
			"marketClose": &graphql.Field{
				Type: graphql.Float,
			},
			"marketHigh": &graphql.Field{
				Type: graphql.Float,
			},
			"marketLow": &graphql.Field{
				Type: graphql.Float,
			},
			"marketVolume": &graphql.Field{
				Type: graphql.Int,
			},
			"marketChangeOverTime": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var Chart = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Chart",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: graphql.NewList(chartData),
			},
		},
	},
)
