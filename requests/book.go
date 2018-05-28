package requests

import (
	"fmt"

	"github.com/diegofu/iex-graphql/types"
)

// GetBook returns a book of a symbol
func GetBook(symbol string) (book types.Book, err error) {
	if err = get(&book, fmt.Sprintf("https://api.iextrading.com/1.0/stock/%s/book", symbol)); err != nil {
		return book, err
	}

	return book, nil
}
