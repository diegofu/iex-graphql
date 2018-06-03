package types

// Financials is a list of income statement, balance sheet, and cash flow data
// from the four most recent reported quarters for a stock
type Financials struct {
	Symbol     string `json:"symbol"`
	Financials []struct {
		ReportDate             string `json:"reportDate"`
		GrossProfit            int64  `json:"grossProfit"`
		CostOfRevenue          int64  `json:"costOfRevenue"`
		OperatingRevenue       int64  `json:"operatingRevenue"`
		TotalRevenue           int64  `json:"totalRevenue"`
		OperatingIncome        int64  `json:"operatingIncome"`
		NetIncome              int64  `json:"netIncome"`
		ResearchAndDevelopment int64  `json:"researchAndDevelopment"`
		OperatingExpense       int64  `json:"operatingExpense"`
		CurrentAssets          int64  `json:"currentAssets"`
		TotalAssets            int64  `json:"totalAssets"`
		TotalLiabilities       int64  `json:"totalLiabilities"`
		CurrentCash            int64  `json:"currentCash"`
		CurrentDebt            int64  `json:"currentDebt"`
		TotalCash              int64  `json:"totalCash"`
		TotalDebt              int64  `json:"totalDebt"`
		ShareholderEquity      int64  `json:"shareholderEquity"`
		CashChange             int64  `json:"cashChange"`
		CashFlow               int64  `json:"cashFlow"`
		OperatingGainsLosses   string `json:"operatingGainsLosses"`
	} `json:"financials"`
}
