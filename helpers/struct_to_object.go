package helpers

import (
	"reflect"
	"unicode"

	"github.com/graphql-go/graphql"
)

// Transform transforms an interface to graphql.Fields
func Transform(s interface{}) *graphql.Object {
	fields := graphql.Fields{}
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldName := field.Tag.Get("json")
		switch field.Type.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			fields[fieldName] = &graphql.Field{
				Type: graphql.Int,
			}
		case reflect.String:
			fields[fieldName] = &graphql.Field{
				Type: graphql.String,
			}
		case reflect.Float32, reflect.Float64:
			fields[fieldName] = &graphql.Field{
				Type: graphql.Float,
			}
		}
	}
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   reflect.TypeOf(s).Elem().Name(),
			Fields: fields,
		},
	)
}

func lowercaseFirstLetter(s string) string {
	letters := []rune(s)
	letters[0] = unicode.ToLower(letters[0])
	return string(letters)
}
