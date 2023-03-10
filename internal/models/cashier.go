package models

// CashType bank note type
type CashType uint

const (
	// CashType1000 1000 baht
	CashType1000 CashType = iota + 1
	// CashType500 500 baht
	CashType500
	// CashType100 100 baht
	CashType100
	// CashType50 50 baht
	CashType50
	// CashType20 20 baht
	CashType20
	// CashType10 10 baht
	CashType10
	// CashType5 5 baht
	CashType5
	// CashType1 1 baht
	CashType1
	// CashTypePoint25 0.25 baht
	CashTypePoint25
)

// CashAmount cash amount
type CashAmount struct {
	CashType CashType `json:"cash_type"`
	Amount   int      `json:"amount"`
}

// GetCash get cash
func (t CashType) GetCash() float64 {
	cash := 0.0
	switch t {
	case CashType1:
		cash = 1
	case CashType10:
		cash = 10
	case CashType100:
		cash = 100
	case CashType1000:
		cash = 1000
	case CashType20:
		cash = 20
	case CashType5:
		cash = 5
	case CashType50:
		cash = 50
	case CashType500:
		cash = 500
	case CashTypePoint25:
		cash = 0.25
	}

	return cash
}

// Type type
func (t CashType) Type() *Type {
	name := "-"
	switch t {
	case CashType1:
		name = "1"
	case CashType10:
		name = "10"
	case CashType100:
		name = "100"
	case CashType1000:
		name = "1000"
	case CashType20:
		name = "20"
	case CashType5:
		name = "5"
	case CashType50:
		name = "50"
	case CashType500:
		name = "500"
	case CashTypePoint25:
		name = "0.25"
	}

	return &Type{
		Name:  name,
		Value: int(t),
	}
}
