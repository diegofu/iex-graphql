package types

type Bid struct {
	Price     float32 `json:"price"`
	Size      int64   `json:"size"`
	Timestamp int64   `json:"timestamp"`
}
