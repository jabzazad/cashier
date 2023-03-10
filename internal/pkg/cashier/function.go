package cashier

import (
	"cashier-api/internal/models"
	"math"
)

// FindUsedAmount find used amount
func (s *service) FindUsedAmount(price float64, cash *models.CashAmount) int {
	usedAmount := math.Floor(price / cash.CashType.GetCash())
	if usedAmount >= float64(cash.Amount) {
		return int(usedAmount)
	}

	return int(usedAmount)
}
