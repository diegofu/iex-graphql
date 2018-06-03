package helpers

import "github.com/graphql-go/graphql"

var quote = graphql.NewObject(
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
			"sector": &graphql.Field{
				Type: graphql.String,
			},
			"calculationPrice": &graphql.Field{
				Type: graphql.String,
			},
			"open": &graphql.Field{
				Type: graphql.Float,
			},
			"openTime": &graphql.Field{
				Type: graphql.Int,
			},
			"close": &graphql.Field{
				Type: graphql.Float,
			},
			"closeTime": &graphql.Field{
				Type: graphql.Int,
			},
			"high": &graphql.Field{
				Type: graphql.Float,
			},
			"low": &graphql.Field{
				Type: graphql.Float,
			},
			"latestPrice": &graphql.Field{
				Type: graphql.Float,
			},
			"latestSource": &graphql.Field{
				Type: graphql.String,
			},
			"latestTime": &graphql.Field{
				Type: graphql.String,
			},
			"latestUpdate": &graphql.Field{
				Type: graphql.Int,
			},
			"latestVolume": &graphql.Field{
				Type: graphql.Int,
			},
			"iexRealtimePrice": &graphql.Field{
				Type: graphql.Float,
			},
			"iexRealtimeSize": &graphql.Field{
				Type: graphql.Int,
			},
			"iexLastUpdated": &graphql.Field{
				Type: graphql.Int,
			},
			"delayedPrice": &graphql.Field{
				Type: graphql.Float,
			},
			"delayedPriceTime": &graphql.Field{
				Type: graphql.Int,
			},
			"extendedPrice": &graphql.Field{
				Type: graphql.Float,
			},
			"extendedChange": &graphql.Field{
				Type: graphql.Float,
			},
			"extendedChangePercent": &graphql.Field{
				Type: graphql.Float,
			},
			"extendedPriceTime": &graphql.Field{
				Type: graphql.Int,
			},
			"change": &graphql.Field{
				Type: graphql.Float,
			},
			"changePercent": &graphql.Field{
				Type: graphql.Float,
			},
			"iexMarketPercent": &graphql.Field{
				Type: graphql.Float,
			},
			"iexVolume": &graphql.Field{
				Type: graphql.Int,
			},
			"avgTotalVolume": &graphql.Field{
				Type: graphql.Int,
			},
			"iexBidPrice": &graphql.Field{
				Type: graphql.Float,
			},
			"iexBidSize": &graphql.Field{
				Type: graphql.Int,
			},
			"iexAskPrice": &graphql.Field{
				Type: graphql.Float,
			},
			"iexAskSize": &graphql.Field{
				Type: graphql.Int,
			},
			"marketCap": &graphql.Field{
				Type: graphql.Int,
			},
			"peRatio": &graphql.Field{
				Type: graphql.Float,
			},
			"week52High": &graphql.Field{
				Type: graphql.Float,
			},
			"week52Low": &graphql.Field{
				Type: graphql.Float,
			},
			"ytdChange": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
