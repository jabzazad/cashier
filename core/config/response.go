package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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
	DuplicateError              Error `mapstructure:"duplicate_error"`
	InvalidParameters           Error `mapstructure:"invalid_parameters"`
	InvalidID                   Error `mapstructure:"invalid_id"`
	DataNotFound                Error `mapstructure:"data_notfound"`
	InvalidLoginType            Error `mapstructure:"invalid_login_type"`
	InvalidEmail                Error `mapstructure:"invalid_email"`
	InvalidPassword             Error `mapstructure:"invalid_password"`
	InvalidFacebookToken        Error `mapstructure:"invalid_facebook_token"`
	InvalidGoogleToken          Error `mapstructure:"invalid_google_token"`
	InvalidRequest              Error `mapstructure:"invalid_request"`
	InvalidPermission           Error `mapstructure:"invalid_permission"`
	InvalidRole                 Error `mapstructure:"invalid_role"`
	AlreadyPhoneNumber          Error `mapstructure:"already_phone_number"`
	AlreadyIDCard               Error `mapstructure:"already_id_card"`
	InvalidIDCard               Error `mapstructure:"invalid_id_card"`
	InvalidPhoneNumber          Error `mapstructure:"invalid_phone_number"`
	InvalidPronoun              Error `mapstructure:"invalid_pronoun"`
	InvalidFirstNameAllSpacebar Error `mapstructure:"invalid_first_name_all_spacebar"`
	InvalidLastNameAllSpacebar  Error `mapstructure:"invalid_last_name_all_spacebar"`
	InvalidAmountPassword       Error `mapstructure:"invalid_amount_password"`
	InvalidIDCardOrPhoneNumber  Error `mapstructure:"invalid_id_card_or_phone_number"`
	PasswordNotMatch            Error `mapstructure:"password_not_match"`
	InvalidPlotName             Error `mapstructure:"invalid_plot_name"`
	InvalidUsernameOrPass       Error `mapstructure:"invalid_username_or_password"`
	InvalidStartAndEndWeek      Error `mapstructure:"invalid_start_end_week"`
	SearchNotFound              Error `mapstructure:"search_not_found"`
	InvalidLoginRole            Error `mapstructure:"invalid_login_role"`
	PermissionNotFound          Error `mapstructure:"permission_not_found"`
	PermissionDenied            Error `mapstructure:"permission_denied"`
	AlreadyEmail                Error `mapstructure:"already_email"`
	InvalidOTP                  Error `mapstructure:"invalid_otp"`
	EmailOrPasswordIncorrect    Error `mapstructure:"email_or_password_incorrect"`
	CannotChangeYourSelfRole    Error `mapstructure:"connot_change_yourself_role"`
	InvalidOldPassword          Error `mapstructure:"invalid_old_password"`
	AlreadyInOrganaization      Error `mapstructure:"already_in_organization"`
	HasOneAdmin                 Error `mapstructure:"has_only_admin"`
	InvalidSMS                  Error `mapstructure:"invalid_sms"`
	InsufficientMoney           Error `mapstructure:"insufficient_money"`
	Internal                    struct {
		Success           Error `mapstructure:"success"`
		UploadSuccess     Error `mapstructure:"upload_success"`
		General           Error `mapstructure:"general"`
		BadRequest        Error `mapstructure:"bad_request"`
		ConnectionError   Error `mapstructure:"connection_error"`
		DBSessionNotFound Error `mapstructure:"db_session_not_found"`
		Unauthorized      Error `mapstructure:"unauthorized"`
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
	switch true {
	case err == bcrypt.ErrMismatchedHashAndPassword:
		return r.InvalidPassword
	default:
		return Error{
			Code: 0,
			Message: localizationMessage{
				EN: err.Error(),
				TH: err.Error(),
			},
		}
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
