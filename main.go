package main

import (
	"fmt"

	"github.com/diegofu/iex-graphql/schemas"
	"github.com/graphql-go/graphql"
)

func main() {
	schema, err := schemas.NewChart()

	if err != nil {
		panic(err)
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: "{chart(range: ONE_DAY, symbol: \"aapl\", date: \"\"){data{minute}, range}}",
	})

	fmt.Printf("%v", result)

}
