package middleware

import (
	"github.com/labstack/echo/v4"
)

// AcceptLanguage accept language
func AcceptLanguage() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lang := c.Request().Header.Get("Accept-Language")
			if lang != "en" {
				if lang != "th" {
					lang = "en"
				}
			}

			c.Set("language", lang)
			return next(c)
		}
	}
}
