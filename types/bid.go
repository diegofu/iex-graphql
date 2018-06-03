package types

// Bid is the bid price of a symbol
type Bid struct {
	Price     float32 `json:"price"`
	Size      int64   `json:"size"`
	Timestamp int64   `json:"timestamp"`
}
