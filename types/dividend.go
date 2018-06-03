package types

// Dividend represents the divident of a symbol
type Dividend struct {
	ExDate       string  `json:"exDate"`
	PaymentDate  string  `json:"paymentDate"`
	RecordDate   string  `json:"recordDate"`
	DeclaredDate string  `json:"declaredDate"`
	Amount       float32 `json:"amount"`
	Flag         string  `json:"flag"`
	Type         string  `json:"type"`
	Qualified    string  `json:"qualified"`
	Indicated    float32 `json:"indicated"`
}
