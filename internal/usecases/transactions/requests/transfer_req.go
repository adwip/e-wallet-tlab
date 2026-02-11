package requests

import (
	"errors"

	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type TransferReq struct {
	Amount float64 `json:"amount" validate:"required,numeric"`
	To     string  `json:"to" validate:"required"`
	Note   string  `json:"note"`
}

func NewTransferReq(c *echo.Context) (out TransferReq, err error) {
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
