package types

// Book is the book value of a symbol
type Book struct {
	Quote Quote `json:"quote"`
	Bids  []Bid `json:"bids"`
	Asks  []Ask `json:"asks"`
}
