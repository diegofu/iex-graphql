package helpers

import (
	"reflect"
	"testing"

	"github.com/diegofu/iex-graphql/types"
)

func TestTransformQuote(t *testing.T) {
	t.Run("Quote", func(t *testing.T) {
		object := Transform(&types.Quote{})
		for k, v := range quote.Fields() {
			if !reflect.DeepEqual(v, object.Fields()[k]) {
				t.Errorf("Wanted %s: %v, Received %s: %v", k, v, k, object.Fields()[k])
			}
		}
		if !reflect.DeepEqual(object, quote) {
			t.Errorf("Wanted %v, Received %v", quote, object)
		}
	})
	t.Run("Chart", func(t *testing.T) {
		object := Transform(&types.Chart{})

		for k, v := range chart.Fields() {
			if !reflect.DeepEqual(v, object.Fields()[k]) {
				t.Errorf("Wanted %s: %v, Received %s: %v", k, v, k, object.Fields()[k])
			}
		}

		if !reflect.DeepEqual(object, chart) {
			t.Errorf("Wanted %v, Received %v", chart, object)
		}
	})
}
