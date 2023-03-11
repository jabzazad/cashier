package context

import (
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	pathKey            = "path"
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
