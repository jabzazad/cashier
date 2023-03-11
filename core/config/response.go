package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var (
	// RR -> for use to return result model
	RR = &Results{}
)

// Error error
type Error struct {
	Code    int                 `json:"code,omitempty" mapstructure:"code"`
	Message localizationMessage `json:"message,omitempty" mapstructure:"localization"`
}

type localizationMessage struct {
	EN     string `mapstructure:"en"`
	TH     string `mapstructure:"th"`
	Locale string `mapstructure:"success"`
}

// WithLocalization with localization language
func (ec Error) WithLocalization(c echo.Context) Error {
	locale, ok := c.Get("language").(string)
	if !ok {
		ec.Message.Locale = "th"
	}
	ec.Message.Locale = locale

	return ec
}

// MarshalJSON marshal json
func (lm localizationMessage) MarshalJSON() ([]byte, error) {
	if strings.ToLower(lm.Locale) == "th" {
		return json.Marshal(lm.TH)
	}

	return json.Marshal(lm.EN)
}

// UnmarshalJSON unmarshal json
func (lm *localizationMessage) UnmarshalJSON(data []byte) error {
	var res string
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	fmt.Println("Unmarshal")
	lm.EN = res
	lm.Locale = "en"
	return nil
}

// Results return results
type Results struct {
	InsufficientMoney Error `mapstructure:"insufficient_money"`
	Internal          struct {
		Success         Error `mapstructure:"success"`
		General         Error `mapstructure:"general"`
		BadRequest      Error `mapstructure:"bad_request"`
		ConnectionError Error `mapstructure:"connection_error"`
	} `mapstructure:"internal"`
}

// Error return error message
func (ec Error) Error() string {
	if ec.Message.Locale == "th" {
		return ec.Message.TH
	}

	return ec.Message.EN
}

// ErrorCode get error code
func (ec Error) ErrorCode() int {
	return ec.Code
}

// GetResponse get error response
func (r *Results) GetResponse(err error) error {
	if _, ok := err.(*echo.HTTPError); ok {
		return err
	} else if _, ok := err.(Error); ok {
		return err
	}

	return Error{
		Code: 0,
		Message: localizationMessage{
			EN: err.Error(),
			TH: err.Error(),
		},
	}
}

// ReadReturnResult read response
func ReadReturnResult(path, filename string) error {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigType("yml")
	v.SetConfigName(filename)
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&RR); err != nil {
		return err
	}

	return nil
}

// CustomErrorMessage custom error message
func (r *Results) CustomErrorMessage(message string) error {
	return Error{
		Code: 999,
		Message: localizationMessage{
			EN: message,
			TH: message,
		},
	}
}

// HTTPStatusCode http status code
func (r *Error) HTTPStatusCode() int {
	switch r.Code {
	case 0, 200: // success
		return http.StatusOK
	case 404: // not found
		return http.StatusNotFound
	case 401: // unauthorized
		return http.StatusUnauthorized
	}

	return http.StatusBadRequest
}
