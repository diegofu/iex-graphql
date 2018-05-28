package queries

import (
	"fmt"

	"github.com/diegofu/iex-graphql/objects"
	"github.com/diegofu/iex-graphql/requests"
	"github.com/graphql-go/graphql"
)

var chartRange = graphql.NewEnum(graphql.EnumConfig{
	Name:        "ChartRange",
	Description: "Time range of the chart",
	Values: graphql.EnumValueConfigMap{
		"FIVE_YEAR": &graphql.EnumValueConfig{
			Value: requests.FiveYear,
		},
		"TWO_YEAR": &graphql.EnumValueConfig{
			Value: requests.TwoYear,
		},
		"ONE_YEAR": &graphql.EnumValueConfig{
			Value: requests.OneYear,
		},
		"YTD": &graphql.EnumValueConfig{
			Value:       requests.Ytd,
			Description: "Year-to-date",
		},
		"SIX_MONTH": &graphql.EnumValueConfig{
			Value: requests.SixMonth,
		},
		"THREE_MONTH": &graphql.EnumValueConfig{
			Value: requests.ThreeMonth,
		},
		"ONE_MONTH": &graphql.EnumValueConfig{
			Value: requests.OneMonth,
		},
		"ONE_DAY": &graphql.EnumValueConfig{
			Value: requests.OneDay,
		},
		"DATE": &graphql.EnumValueConfig{
			Value: requests.Date,
		},
		"DYNAMIC": &graphql.EnumValueConfig{
			Value: requests.Dynamic,
			Description: `Will return 3d or 1m data
			depending on the day or week and time of day.
			Intraday per minute data is only returned during market hours.`,
		},
	},
})

var Chart = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ChartQuery",
		Fields: graphql.Fields{
			"chart": &graphql.Field{
				Type: objects.Chart,
				Args: graphql.FieldConfigArgument{
					"range": &graphql.ArgumentConfig{
						Type: chartRange,
					},
					"symbol": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					timeRange, ok := p.Args["range"].(requests.Range)
					if !ok {
						return nil, fmt.Errorf("Unable to parse range %v", p.Args["range"])
					}
					symbol, ok := p.Args["symbol"].(string)
					if !ok {
						return nil, fmt.Errorf("Unable to parse symbol %v", p.Args["symbol"])
					}
					date, ok := p.Args["date"].(string)
					if !ok {
						return nil, fmt.Errorf("Unable to parse date %v", p.Args["date"])
					}
					chart, err := requests.GetChart(symbol, timeRange, date)

					if err != nil {
						return nil, err
					}

					return chart, nil
				},
			},
		},
	},
)
