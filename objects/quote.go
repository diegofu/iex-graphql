package objects

import (
	"github.com/diegofu/iex-graphql/helpers"
	"github.com/diegofu/iex-graphql/types"
	"github.com/graphql-go/graphql"
)

var calculationPrice = graphql.NewEnum(graphql.EnumConfig{
	Name:        "CalculationPrice",
	Description: "refers to the source of the latest price.",
	Values: graphql.EnumValueConfigMap{
		"TOPS": &graphql.EnumValueConfig{
			Value: "tops",
		},
		"SIP": &graphql.EnumValueConfig{
			Value: "sip",
		},
		"PREVIOUS_CLOSE": &graphql.EnumValueConfig{
			Value: "previousclose",
		},
		"CLOSE": &graphql.EnumValueConfig{
			Value: "close",
		},
	},
})

// Quote is the graphql representation of a quote
var Quote = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Quote",
		Fields: helpers.Transform(&types.Quote{}),
	},
)
