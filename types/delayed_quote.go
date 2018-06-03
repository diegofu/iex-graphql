package types

// DelayedQuote is the 15 minute delayed market quote
type DelayedQuote struct {
	Symbol           string  `json:"symbol"`
	DelayedPrice     float32 `json:"delayedPrice"`
	DelayedSize      int64   `json:"delayedSize"`
	DelayedPriceTime int64   `json:"delayedPriceTime"`
	ProcessedTime    int64   `json:"processedTime"`
}
