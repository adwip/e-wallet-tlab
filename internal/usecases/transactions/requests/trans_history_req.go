package requests

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type TransactionHistoryReq struct {
	Limit  int `query:"limit" validate:"required,numeric,min=10"`
	Offset int `query:"offset"`
}

func NewTransactionHistoryReq(c *echo.Context) (out TransactionHistoryReq, err error) {
	if err = c.Bind(&out); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if err = c.Validate(&out); err != nil {
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	return out, nil
}
