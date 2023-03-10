package request

import "cashier-api/internal/models"

// PaymentRequest payment request
type PaymentRequest struct {
	ProductPrice float64 `json:"product_price"`
	Receiveds    []struct {
		Amount   int             `json:"amount"`
		CashType models.CashType `json:"cash_type"`
	} `json:"receiveds"`
}

type AddMoneyRequest struct {
	Receiveds []struct {
		Amount   int             `json:"amount"`
		CashType models.CashType `json:"cash_type"`
	} `json:"receiveds"`
}
