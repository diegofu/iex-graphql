package types

type Book struct {
	Quote Quote `json:"quote"`
	Bids  []Bid `json:"bids"`
	Asks  []Ask `json:"asks"`
}
