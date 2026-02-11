package requests

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
	"github.com/microcosm-cc/bluemonday"
)

type UserLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func NewUserLoginReq(in *echo.Context) (out UserLoginReq, err error) {
	if err = in.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if err = in.Validate(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	p := bluemonday.StrictPolicy()
	out.Email = p.Sanitize(out.Email)
	return out, nil
}
