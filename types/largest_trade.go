package types

// LargestTrade is struct of a trade
type LargestTrade struct {
	Price     float32 `json:"price"`
	Size      int64   `json:"size"`
	Time      int64   `json:"time"`
	TimeLabel string  `json:"timeLabel"`
	Venue     string  `json:"venue"`
	VenueName string  `json:"venueName"`
}
