package helpers

import "github.com/graphql-go/graphql"

var chart = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Chart",
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
			"average": &graphql.Field{
				Type: graphql.Float,
			},
			"notional": &graphql.Field{
				Type: graphql.Float,
			},
			"numberOfTrades": &graphql.Field{
				Type: graphql.Int,
			},
			"high": &graphql.Field{
				Type: graphql.Float,
			},
			"low": &graphql.Field{
				Type: graphql.Float,
			},
			"volume": &graphql.Field{
				Type: graphql.Int,
			},
			"label": &graphql.Field{
				Type: graphql.String,
			},
			"changeOverTime": &graphql.Field{
				Type: graphql.Float,
			},
			"date": &graphql.Field{
				Type: graphql.String,
			},
			"open": &graphql.Field{
				Type: graphql.Float,
			},
			"close": &graphql.Field{
				Type: graphql.Float,
			},
			"unadjustedVolume": &graphql.Field{
				Type: graphql.Int,
			},
			"change": &graphql.Field{
				Type: graphql.Float,
			},
			"changePercent": &graphql.Field{
				Type: graphql.Float,
			},
			"vwap": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
