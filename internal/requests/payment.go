package request

// PaymentRequest payment request
type PaymentRequest struct {
	ProductPrice float64 `json:"product_price"`
	Receives     []struct {
		Amount    int     `json:"amount"`
		CashValue float64 `json:"cash_value"`
	} `json:"receives"`
}

// AddMoneyRequest add money request
type AddMoneyRequest struct {
	Receiveds []struct {
		Amount    int     `json:"amount"`
		CashValue float64 `json:"cash_value"`
	} `json:"receiveds"`
}

// UpdateCashNoteRequest update cash note request
type UpdateCashNoteRequest struct {
	CashValue float64 `json:"cash_value" path:"cash_value"`
	Amount    int     `json:"amount"`
}
