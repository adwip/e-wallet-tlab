package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type FormValidator struct {
	Validator *validator.Validate
}

func (cv *FormValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}
