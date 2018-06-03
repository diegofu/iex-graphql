package types

// Split is a split of a stock
type Split struct {
	ExDate       string  `json:"exDate"`       // refers to the split ex-date
	DeclaredDate string  `json:"declaredDate"` // refers to the split declaration date
	RecordDate   string  `json:"recordDate"`   // refers to the split record date
	PaymentDate  string  `json:"paymentDate"`  // refers to the split payment date
	Ratio        float32 `json:"ratio"`        // The split ratio is an inverse of the number of shares that a holder of the stock would have after the split divided by the number of shares that the holder had before. For example: Split ratio of .5 = 2 for 1 split.
	ToFactor     int64   `json:"toFactor"`     // To factor of the split. Used to calculate the split ratio forfactor/tofactor = ratio (eg ½ = 0.5)
	ForFactor    int64   `json:"forFactor"`    // For factor of the split. Used to calculate the split ratio forfactor/tofactor = ratio (eg ½ = 0.5)
}
