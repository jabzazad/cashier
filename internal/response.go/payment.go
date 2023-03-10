package response

import "cashier-api/internal/models"

// ChangeResponse change response
type ChangesResponse struct {
	Changes []*ChangeResponse `json:"changes"`
}

type ChangeResponse struct {
	Amount   int             `json:"amount"`
	CashType models.CashType `json:"cash_type"`
}
