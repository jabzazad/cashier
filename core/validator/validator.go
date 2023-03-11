package validator

import (
	"cashier-api/core/translator"

	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// APValidator validator
type APValidator struct {
	validator *validator.Validate
}

// New new
func New() *APValidator {
	v := &APValidator{
		validator: validator.New(),
	}
	v.translator()
	return v
}

func (cv *APValidator) translator() {
	if err := en_translations.RegisterDefaultTranslations(cv.validator, translator.ENTranslator); err != nil {
		panic(err)
	}
	_ = cv.validator.RegisterTranslation("required", translator.ENTranslator,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} is a required field", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)
}

// Validate validator
func (cv *APValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
