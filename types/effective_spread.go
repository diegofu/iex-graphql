package types

// EffectiveSpread represents effective spread, eligible volume,
// and price improvement of a stock, by market
type EffectiveSpread struct {
	Volume           int64   `json:"volume"`
	Venue            string  `json:"venue"`
	VenueName        string  `json:"venueName"`
	EffectiveSpread  float32 `json:"effectiveSpread"`
	EffectiveQuoted  float32 `json:"effectiveQuoted"`
	PriceImprovement float32 `json:"priceImprovement"`
}
