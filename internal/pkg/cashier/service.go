package cashier

import (
	"cashier-api/core/config"
	"cashier-api/core/context"
	"cashier-api/core/logger"
	"cashier-api/internal/models"
	request "cashier-api/internal/requests"
	"cashier-api/internal/response.go"
	"cashier-api/internal/utils"
	"sync"
)

// Service service interface
type Service interface {
	GetCashierDesk(c *context.Context) ([]*models.CashAmount, error)
	CalculateChange(c *context.Context, request *request.PaymentRequest) (*response.ChangesResponse, error)
	AddCash(c *context.Context, request *request.AddMoneyRequest) error
	UpdateMoneyNoteAmount(c *context.Context, request *request.UpdateCashNoteRequest) error
}

type service struct {
	env   *config.Environment
	rr    *config.Results
	mutex sync.Mutex
}

// NewService new service
func NewService() Service {
	return &service{
		env: config.ENV,
		rr:  config.RR,
	}
}

// GetCashierDesk get cashier desk
func (s *service) GetCashierDesk(c *context.Context) ([]*models.CashAmount, error) {
	entities := []*models.CashAmount{}
	err := utils.ReadJSONFile(s.env.CashierDeskPath, &entities)
	if err != nil {
		logger.Logger.Errorf("get cashier desk data error: %s", err)
		return nil, err
	}

	return entities, nil
}

// CalculateChange calculate change
func (s *service) CalculateChange(c *context.Context, request *request.PaymentRequest) (*response.ChangesResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var totalReceivedCash float64
	cashDesk, err := s.GetCashierDesk(c)
	if err != nil {
		return nil, err
	}

	for _, cash := range request.Receiveds {
		totalReceivedCash += (float64(cash.Amount) * cash.CashType.GetCash())
		for _, cashierCash := range cashDesk {
			if cashierCash.CashType == cash.CashType {
				cashierCash.Amount += cash.Amount
			}
		}
	}

	changeResponse := &response.ChangesResponse{}
	totalChange := totalReceivedCash - request.ProductPrice
	for _, cash := range cashDesk {
		used := s.FindUsedAmount(totalChange, cash)
		if used > 0 {
			totalChange -= (float64(used) * cash.CashType.GetCash())
			cash.Amount -= used
			changeResponse.Changes = append(changeResponse.Changes, &response.ChangeResponse{
				CashType: cash.CashType,
				Amount:   used,
			})
		}
	}

	if totalChange > 0 {
		return nil, s.rr.InsufficientMoney
	}

	err = utils.WriteJsonFile(s.env.CashierDeskPath, cashDesk)
	if err != nil {
		return nil, err
	}

	return changeResponse, nil
}

// AddCash add cash
func (s *service) AddCash(c *context.Context, request *request.AddMoneyRequest) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	cashDesk, err := s.GetCashierDesk(c)
	if err != nil {
		return err
	}

	for _, cash := range cashDesk {
		for _, requestCash := range request.Receiveds {
			if cash.CashType == requestCash.CashType {
				cash.Amount += requestCash.Amount
			}
		}
	}

	err = utils.WriteJsonFile(s.env.CashierDeskPath, cashDesk)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMoneyNoteAmount update money note amount
func (s *service) UpdateMoneyNoteAmount(c *context.Context, request *request.UpdateCashNoteRequest) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	cashDesk, err := s.GetCashierDesk(c)
	if err != nil {
		return err
	}

	for _, cash := range cashDesk {
		if cash.CashType == request.ID {
			cash.Amount = request.Amount
		}
	}

	err = utils.WriteJsonFile(s.env.CashierDeskPath, cashDesk)
	if err != nil {
		return err
	}

	return nil
}
