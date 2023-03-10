package middleware

import (
	"cashier-api/core/context"

	"github.com/labstack/echo/v4"
)

// CustomContext custom context
func CustomContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.Context{Context: c}
			return next(cc)
		}
	}
}
