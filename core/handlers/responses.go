package handlers

import (
	"cashier-api/core/context"
	"cashier-api/core/logger"
	"cashier-api/core/render"
	"cashier-api/internal/models"
	"reflect"

	"github.com/labstack/echo/v4"
)

// ResponseObject handle response object
func ResponseObject(c echo.Context, fn interface{}, request interface{}) error {
	ctx := c.(*context.Context)
	err := ctx.BindAndValidate(request)
	if err != nil {
		logger.Logger.Errorf("bind value error: %s", err)
		return render.Error(c, err)
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(request),
	})
	errObj := out[1].Interface()
	if errObj != nil {
		logger.Logger.Errorf("call service error: %s", errObj)
		return render.Error(c, errObj.(error))
	}

	return render.JSON(c, out[0].Interface())
}

// ResponseOnlyObject handle response object
func ResponseOnlyObject(c echo.Context, fn interface{}, request interface{}) error {
	ctx := c.(*context.Context)
	err := ctx.BindAndValidate(request)
	if err != nil {
		logger.Logger.Errorf("bind value error: %s", err)
		return render.Error(c, err)
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(request),
	})

	return render.JSON(c, out[0].Interface())
}

// ResponseObjectWithoutRequest handle response object without request
func ResponseObjectWithoutRequest(c echo.Context, fn interface{}) error {
	ctx := c.(*context.Context)
	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(ctx),
	})
	errObj := out[1].Interface()
	if errObj != nil {
		logger.Logger.Errorf("call service error: %s", errObj)
		return render.Error(c, errObj.(error))
	}

	return render.JSON(c, out[0].Interface())
}

// ResponseSuccess handle response success
func ResponseSuccess(c echo.Context, fn interface{}, request interface{}) error {
	ctx := c.(*context.Context)
	err := ctx.BindAndValidate(request)
	if err != nil {
		logger.Logger.Errorf("bind value error: %s", err)
		return render.Error(c, err)
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(request),
	})
	errObj := out[0].Interface()
	if errObj != nil {
		logger.Logger.Errorf("call service error: %s", errObj)
		return render.Error(c, errObj.(error))
	}
	return render.JSON(c, models.NewSuccessMessage())
}

// ResponseSuccessWithoutRequest handle response success without request
func ResponseSuccessWithoutRequest(c echo.Context, fn interface{}) error {
	ctx := c.(*context.Context)
	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(ctx),
	})
	errObj := out[0].Interface()
	if errObj != nil {
		logger.Logger.Errorf("call service error: %s", errObj)
		return render.Error(c, errObj.(error))
	}
	return render.JSON(c, models.NewSuccessMessage())
}
