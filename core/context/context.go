package context

import (
	"reflect"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	pathKey            = "path"
	databaseKey        = "sql"
	compositeFormDepth = 2
)

// Context custom echo context
type Context struct {
	echo.Context
	Parameters interface{}
}

// BindAndValidate bind and validate form
func (c *Context) BindAndValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	c.parsePathParams(i, 1)
	if err := c.Validate(i); err != nil {
		return err
	}

	c.Parameters = i
	return nil
}

// SetDatabase set database
func (c *Context) SetDatabase(database *gorm.DB) {
	c.Set(databaseKey, database)
}

// GetDatabase get database
func (c *Context) GetDatabase() (*gorm.DB, bool) {
	databaseConnection := c.Get(databaseKey)
	if databaseConnection == nil {
		return nil, false
	}
	return databaseConnection.(*gorm.DB), true
}

func (c *Context) parsePathParams(form interface{}, depth int) {
	formValue := reflect.ValueOf(form)
	if formValue.Kind() == reflect.Ptr {
		formValue = formValue.Elem()
	}

	t := reflect.TypeOf(formValue.Interface())
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		paramValue := formValue.FieldByName(fieldName)
		if paramValue.IsValid() {
			if depth < compositeFormDepth && paramValue.Kind() == reflect.Struct {
				depth++
				c.parsePathParams(paramValue.Addr().Interface(), depth)
			}
			tag := t.Field(i).Tag.Get(pathKey)
			if tag != "" {
				value := c.Param(tag)
				if paramValue.Kind() == reflect.Uint {
					number, _ := strconv.ParseUint(value, 10, 64)
					paramValue.SetUint(number)
					continue
				}

				paramValue.Set(reflect.ValueOf(c.Param(tag)))

			}
		}

	}
}

// Claims jwt claims
type Claims struct {
	jwt.StandardClaims
	UserID         uint `json:"user_id"`
	RefreshTokenID uint `json:"refresh_token_id"`
}

// GetUserSession get user session
func (c *Context) GetUserSession() *UserContext {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*Claims).ToUserSession()
}

// ToUserSession convert claims to user session
func (c *Claims) ToUserSession() *UserContext {
	return &UserContext{
		UserID:         c.UserID,
		RefreshTokenID: c.RefreshTokenID,
	}
}

// UserContext user context
type UserContext struct {
	UserID         uint
	RefreshTokenID uint
}
