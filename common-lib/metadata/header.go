package metadata

import (
	"context"

	"github.com/adwip/e-wallet-tlab/common-lib/shared/constant"
	"github.com/labstack/echo/v5"
)

func GetRpcRequestId(ctx context.Context) string {
	header, isset := ctx.Value(constant.ContextKey).(map[string]string)
	if !isset {
		return ""
	}
	return header[XRequestId]
}

func GetRpcAuthUserId(ctx context.Context) string {
	header, isset := ctx.Value(constant.ContextKey).(map[string]string)
	if !isset {
		return ""
	}
	return header[XAuthUserId]
}

func GetRequestId(ctx *echo.Context) string {
	header := ctx.Request().Header.Get(XRequestId)
	return header
}

func GetAuthUserId(ctx *echo.Context) string {
	header := ctx.Request().Header.Get(XAuthUserId)
	return header
}
