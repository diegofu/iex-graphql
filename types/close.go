package types

// Close represents price for when close
type Close struct {
	Price float32 `json:"price"`
	Time  int64   `json:"time"`
}
