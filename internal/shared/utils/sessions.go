package utils

import (
	"github.com/labstack/echo/v5"
)

func GetUserId(ctx *echo.Context) string {
	return ctx.Get("user_id").(string)
}

func GetUserName(ctx *echo.Context) string {
	return ctx.Get("user_name").(string)
}
