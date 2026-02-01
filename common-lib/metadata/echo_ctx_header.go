package metadata

import "github.com/labstack/echo/v5"

func ReadXAuthUserId(ctx *echo.Context) string {
	return ctx.Request().Header.Get(XAuthUserId)
}

func ReadXRequestId(ctx *echo.Context) string {
	return ctx.Request().Header.Get(XRequestId)
}
