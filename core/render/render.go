// Package render is a internal handlers render package
package render

import (
	"cashier-api/core/config"

	"github.com/labstack/echo/v4"
)

// JSON render json to client
func JSON(c echo.Context, response interface{}) error {
	return c.
		JSON(config.RR.Internal.Success.HTTPStatusCode(), response)
}

// Error render error to client
func Error(c echo.Context, err error) error {
	errMsg := config.RR.Internal.ConnectionError
	if locErr, ok := err.(config.Error); ok {
		errMsg = locErr
	}

	return c.
		JSON(errMsg.HTTPStatusCode(), errMsg)
}
