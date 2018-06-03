package types

// Company represents the detail fo a symbol
type Company struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
	Exchange    string `json:"exchange"`
	Industry    string `json:"industry"`
	Website     string `json:"website"`
	Description string `json:"description"`
	CEO         string `json:"CEO"`
	IssueType   string `json:"issueType"`
	Sector      string `json:"sector"`
}
