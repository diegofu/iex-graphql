package types

// Open represents price for when close
type Open struct {
	Price float32 `json:"price"`
	Time  int64   `json:"time"`
}
