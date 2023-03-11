package models

// CashType bank note type
type CashType uint

// CashAmount cash amount
type CashAmount struct {
	CashValue float64 `json:"cash_value "`
	Amount    int     `json:"amount"`
}
