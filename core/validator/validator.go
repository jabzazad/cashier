package validator

import (
	"cashier-api/core/translator"
	"unicode"

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
	v.customValidateor()
	v.translator()
	return v
}

func (cv *APValidator) customValidateor() {
	_ = cv.validator.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return validPassword(fl.Field().String())
	})
}

func validPassword(s string) bool {
	/*
	 * Password rules:
	 * at least 8 letters
	 * at least 1 number
	 * at least 1 upper case
	 */
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(s) >= 8 {
		hasMinLen = true
	}
	for _, char := range s {

		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	valid := hasMinLen && hasUpper && hasLower && hasNumber
	return valid
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
	cv.addTranslation("password", "รหัสผ่านเดิมของคุณไม่ถูกรูปแบบ / Your password invalid format")
}

func (cv *APValidator) addTranslation(name, message string) {
	_ = cv.validator.RegisterTranslation(name, translator.ENTranslator,
		func(ut ut.Translator) error {
			return ut.Add(name, message, true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(name, fe.Field())
			return t
		},
	)
}

// Validate validator
func (cv *APValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Var var
func (cv *APValidator) Var(field interface{}, tag string) error {
	return cv.validator.Var(field, tag)
}
