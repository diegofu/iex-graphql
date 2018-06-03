package types

// KeyStats contains all the key stats of a stock
type KeyStats struct {
	CompanyName         string  `json:"companyName"`
	Marketcap           int64   `json:"marketcap"`
	Beta                float32 `json:"beta"`
	Week52high          float32 `json:"week52high"`
	Week52low           float32 `json:"week52low"`
	Week52change        float32 `json:"week52change"`
	ShortInterest       float32 `json:"shortInterest"`
	ShortDate           string  `json:"shortDate"`
	DividendRate        float32 `json:"dividendRate"`
	DividendYield       float32 `json:"dividendYield"`
	ExDividendDate      string  `json:"exDividendDate"`
	LatestEPS           float32 `json:"latestEPS"`
	LatestEPSDate       string  `json:"latestEPSDate"`
	SharesOutstanding   int64   `json:"sharesOutstanding"`
	Float               int64   `json:"float"`
	ReturnOnEquity      float32 `json:"returnOnEquity"`
	ConsensusEPS        float32 `json:"consensusEPS"`
	NumberOfEstimates   float32 `json:"float32OfEstimates"`
	Symbol              string  `json:"symbol"`
	EBITDA              int64   `json:"EBITDA"`
	Revenue             int64   `json:"revenue"`
	GrossProfit         int64   `json:"grossProfit"`
	Cash                int64   `json:"cash"`
	Debt                int64   `json:"debt"`
	TtmEPS              float32 `json:"ttmEPS"`
	RevenuePerShare     float32 `json:"revenuePerShare"`
	RevenuePerEmployee  float32 `json:"revenuePerEmployee"`
	PeRatioHigh         float32 `json:"peRatioHigh"`
	PeRatioLow          float32 `json:"peRatioLow"`
	EPSSurpriseDollar   float32 `json:"EPSSurpriseDollar"`
	EPSSurprisePercent  float32 `json:"EPSSurprisePercent"`
	ReturnOnAssets      float32 `json:"returnOnAssets"`
	ReturnOnCapital     float32 `json:"returnOnCapital"`
	ProfitMargin        float32 `json:"profitMargin"`
	PriceToSales        float32 `json:"priceToSales"`
	PriceToBook         float32 `json:"priceToBook"`
	Day200MovingAvg     float32 `json:"day200MovingAvg"`
	Day50MovingAvg      float32 `json:"day50MovingAvg"`
	InstitutionPercent  float32 `json:"institutionPercent"`
	InsiderPercent      float32 `json:"insiderPercent"`
	ShortRatio          float32 `json:"shortRatio"`
	Year5ChangePercent  float32 `json:"year5ChangePercent"`
	Year2ChangePercent  float32 `json:"year2ChangePercent"`
	Year1ChangePercent  float32 `json:"year1ChangePercent"`
	YtdChangePercent    float32 `json:"ytdChangePercent"`
	Month6ChangePercent float32 `json:"month6ChangePercent"`
	Month3ChangePercent float32 `json:"month3ChangePercent"`
	Month1ChangePercent float32 `json:"month1ChangePercent"`
	Day5ChangePercent   float32 `json:"day5ChangePercent"`
}
