package session

import (
	"net/http"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
	"github.com/labstack/echo/v5"
)

type resultFormatter struct {
	Code    string `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func SetResult(ctx *echo.Context, result any, err error) error {
	var (
		httpCode = http.StatusOK
	)
	errCode, _, _, errMsg, _ := stacktrace.DefineStacktrace(err)
	if err != nil {
		httpCode, _, _ = stacktrace.StacktraceToHttpCode(errCode)
	}
	if errMsg == "" && err != nil {
		errMsg = stacktrace.StacktraceMessageByCode(errCode)
	}
	ctx.JSON(httpCode, resultFormatter{
		Code:    errCode,
		Data:    result,
		Message: errMsg,
	})
	return err
}
