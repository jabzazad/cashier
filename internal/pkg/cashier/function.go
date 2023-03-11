package cashier

import (
	"cashier-api/internal/models"
	"math"
)

// FindUsedAmount find used amount
func (s *service) FindUsedAmount(price float64, cash *models.CashAmount) int {
	usedAmount := math.Floor(price / cash.CashValue)
	if usedAmount > float64(cash.Amount) {
		return cash.Amount
	}

	return int(usedAmount)
}
