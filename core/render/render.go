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

// Download render file
func Download(c echo.Context, path, fileName string) error {
	return c.Attachment(path, fileName)
}

// Error render error to client
func Error(c echo.Context, err error) error {
	errMsg := config.RR.Internal.ConnectionError
	if locErr, ok := err.(config.Error); ok {
		errMsg = locErr
	}

	return c.
		JSON(errMsg.HTTPStatusCode(), errMsg.Error())
}
