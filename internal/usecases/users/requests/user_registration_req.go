package requests

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
	"github.com/microcosm-cc/bluemonday"
)

type UserRegistrationReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

func NewUserRegistrationReq(in *echo.Context) (out UserRegistrationReq, err error) {
	if err = in.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if err = in.Validate(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	p := bluemonday.StrictPolicy()
	out.Name = p.Sanitize(out.Name)
	out.Email = p.Sanitize(out.Email)
	out.Password = p.Sanitize(out.Password)
	return out, nil
}
