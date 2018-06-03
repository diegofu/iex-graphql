package types

// Earnings reprenset the four most recent reported quarters
type Earnings struct {
	Symbol   string `json:"symbol"`
	Earnings []struct {
		ActualEPS              float32 `json:"actualEPS"`
		ConsensusEPS           float32 `json:"consensusEPS"`
		EstimatedEPS           float32 `json:"estimatedEPS"`
		AnnounceTime           string  `json:"announceTime"`
		NumberOfEstimates      int64   `json:"numberOfEstimates"`
		EPSSurpriseDollar      float32 `json:"EPSSurpriseDollar"`
		EPSReportDate          string  `json:"EPSReportDate"`
		FiscalPeriod           string  `json:"fiscalPeriod"`
		FiscalEndDate          string  `json:"fiscalEndDate"`
		YearAgo                float32 `json:"yearAgo"`
		YearAgoChangePercent   float32 `json:"yearAgoChangePercent"`
		EstimatedChangePercent float32 `json:"estimatedChangePercent"`
		SymbolID               int     `json:"symbolId"`
	} `json:"earnings"`
}
