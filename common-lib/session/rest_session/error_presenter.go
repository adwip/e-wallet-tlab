package rest_session

import (
	"github.com/adwip/e-wallet-tlab/common-lib/metadata"
	"github.com/labstack/echo/v5"
)

func (s *restSession) ErrorHandler(ctx *echo.Context, err error) {

	s.writeRestLog(err, ctx.Request().Method, ctx.Request().Header.Get(metadata.XRequestId), ctx.Path())
}
