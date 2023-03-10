package middleware

import (
	"cashier-api/core/config"

	"github.com/labstack/echo/v4"
)

// WrapResponse wrap response
func WrapResponse(rr *config.Results) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				return rr.GetResponse(err)
			}
			return nil
		}
	}
}
