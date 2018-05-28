package requests

import (
	"fmt"

	"github.com/diegofu/iex-graphql/types"
)

// Range is a range the chart represents
type Range string

const (
	FiveYear   Range = "5y"
	TwoYear    Range = "2y"
	OneYear    Range = "1y"
	Ytd        Range = "ytd"
	SixMonth   Range = "6m"
	ThreeMonth Range = "3m"
	OneMonth   Range = "1m"
	OneDay     Range = "1d"
	Date       Range = "date"
	Dynamic    Range = "dynamic"
)

// GetChart returns a chart of a certain range
func GetChart(symbol string, r Range, date string) (chartDynamic types.ChartDynamic, err error) {
	if r == Dynamic {
		if err = get(
			&chartDynamic,
			fmt.Sprintf(
				"https://api.iextrading.com/1.0/stock/%s/chart/%s/%s",
				symbol, r, date,
			),
		); err != nil {
			return chartDynamic, err
		}

		return chartDynamic, nil
	}

	chart := new([]types.Chart)
	if err = get(
		chart,
		fmt.Sprintf(
			"https://api.iextrading.com/1.0/stock/%s/chart/%s/%s",
			symbol, r, date,
		),
	); err != nil {
		return chartDynamic, err
	}
	chartDynamic = types.ChartDynamic{
		Data:  *chart,
		Range: string(r),
	}

	return chartDynamic, nil
}
