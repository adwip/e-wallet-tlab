package requests

import (
	"errors"

	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type TopUpReq struct {
	Amount float64 `json:"amount" validate:"required,numeric,min=1"`
}

func NewTopUpReq(c *echo.Context) (out TopUpReq, err error) {
	if err = c.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if err = c.Validate(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if out.Amount <= 0 {
		err = errors.New("amount must be greater than 0")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	return out, nil
}
