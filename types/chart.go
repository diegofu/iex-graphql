package types

type Chart struct {
	Minute               string  `json:"minute"`
	MarketAverage        float32 `json:"marketAverage"`
	MarketNational       float32 `json:"marketNational"`
	MarketNumberOfTrades int64   `json:"marketNumberOfTrades"`
	MarketOpen           float32 `json:"marketOpen"`
	MarketClose          float32 `json:"marketClose"`
	MarketHigh           float32 `json:"marketHigh"`
	MarketLow            float32 `json:"marketLow"`
	MarketVolume         int64   `json:"marketVolume"`
	MarketChangeOverTime float32 `json:"marketChangeOverTime"`
	Average              float32 `json:"average"`
	Notional             float32 `json:"notional"`
	NumberOfTrades       int64   `json:"numberOfTrades"`
	High                 float32 `json:"high"`
	Low                  float32 `json:"low"`
	Volume               int64   `json:"volume"`
	Label                string  `json:"label"`
	ChangeOverTime       float32 `json:"changeOverTime"`
	Date                 string  `json:"date"`
	Open                 float32 `json:"open"`
	Close                float32 `json:"close"`
	UnadjustedVolume     int64   `json:"unadjustedVolume"`
	Change               float32 `json:"change"`
	ChangePercent        float32 `json:"changePercent"`
	Vwap                 float32 `json:"vwap"`
}

type ChartDynamic struct {
	Range string  `json:"range"`
	Data  []Chart `json:"data"`
}
