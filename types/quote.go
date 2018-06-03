package types

// Quote represents the quote of a stock Symbol
type Quote struct {
	Symbol                string  `json:"symbol"`
	CompanyName           string  `json:"companyName"`
	PrimaryExchange       string  `json:"primaryExchange"`
	Sector                string  `json:"sector"`
	CalculationPrice      string  `json:"calculationPrice"`
	Open                  float32 `json:"open"`
	OpenTime              int64   `json:"openTime"`
	Close                 float32 `json:"close"`
	CloseTime             int64   `json:"closeTime"`
	High                  float32 `json:"high"`
	Low                   float32 `json:"low"`
	LatestPrice           float32 `json:"latestPrice"`
	LatestSource          string  `json:"latestSource"`
	LatestTime            string  `json:"latestTime"`
	LatestUpdate          int64   `json:"latestUpdate"`
	LatestVolume          int64   `json:"latestVolume"`
	IEXRealtimePrice      float32 `json:"iexRealtimePrice"`
	IEXRealtimeSize       int64   `json:"iexRealtimeSize"`
	IEXLastUpdated        int64   `json:"iexLastUpdated"`
	DelayedPrice          float32 `json:"delayedPrice"`
	DelayedPriceTime      int64   `json:"delayedPriceTime"`
	ExtendedPrice         float32 `json:"extendedPrice"`
	ExtendedChange        float32 `json:"extendedChange"`
	ExtendedChangePercent float32 `json:"extendedChangePercent"`
	ExtendedPriceTime     int64   `json:"extendedPriceTime"`
	Change                float32 `json:"change"`
	ChangePercent         float32 `json:"changePercent"`
	IEXMarketPercent      float32 `json:"iexMarketPercent"`
	IEXVolume             int64   `json:"iexVolume"`
	AvgTotalVolume        int64   `json:"avgTotalVolume"`
	IEXBidPrice           float32 `json:"iexBidPrice"`
	IEXBidSize            int64   `json:"iexBidSize"`
	IEAskPrice            float32 `json:"iexAskPrice"`
	IEXAskSize            int64   `json:"iexAskSize"`
	MarketCap             int64   `json:"marketCap"`
	PeRatio               float32 `json:"peRatio"`
	Week52High            float32 `json:"week52High"`
	Week52Low             float32 `json:"week52Low"`
	YtdChange             float32 `json:"ytdChange"`
}
