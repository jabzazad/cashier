package response

// ChangeResponse change response
type ChangesResponse struct {
	Changes []*ChangeResponse `json:"changes,omitempty"`
}

type ChangeResponse struct {
	Amount    int     `json:"amount"`
	CashValue float64 `json:"cash_value"`
}
