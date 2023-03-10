package cashier

import (
	"cashier-api/core/config"
	"cashier-api/core/handlers"
	request "cashier-api/internal/requests"

	"github.com/labstack/echo/v4"
)

// Endpoint endpoint interface
type Endpoint interface {
	CalculateChange(c echo.Context) error
	GetCashierDesk(c echo.Context) error
	AddCash(c echo.Context) error
}

type endpoint struct {
	env     *config.Environment
	rr      *config.Results
	service Service
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		env:     config.ENV,
		rr:      config.RR,
		service: NewService(),
	}
}

// CalculateChange calculate change
// @Tags Cashier
// @Summary CalculateChange
// @Description Calculate change
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param payment body request.PaymentRequest true "Payment Request"
// @Success 200 {object} response.ChangesResponse
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /cashier/payments [post]
func (ep *endpoint) CalculateChange(c echo.Context) error {
	return handlers.ResponseObject(c, ep.service.CalculateChange, &request.PaymentRequest{})
}

// GetCashierDesk get cashier desk
// @Tags Cashier
// @Summary GetCashierDesk
// @Description Get cashier desk enum of cash type: {1 : 1000 baht 2 : 500 baht 3 : 100 baht,4 : 50 baht,5 : 20 baht,6 : 10 baht,7 : 5 baht,8 : 1 baht,9 : 0.25 baht }
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {array} models.CashAmount
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /cashier/cash [get]
func (ep *endpoint) GetCashierDesk(c echo.Context) error {
	return handlers.ResponseObjectWithoutRequest(c, ep.service.GetCashierDesk)
}

// AddCash add cash
// @Tags Cashier
// @Summary AddCash
// @Description Add cash
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param cash body request.AddMoneyRequest true "Cash Request"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Router /cashier/cash [post]
func (ep *endpoint) AddCash(c echo.Context) error {
	return handlers.ResponseSuccess(c, ep.service.AddCash, &request.AddMoneyRequest{})
}
