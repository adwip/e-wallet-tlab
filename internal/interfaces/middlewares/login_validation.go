package middlewares

import (
	"strings"

	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/labstack/echo/v5"
)

func LoginValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		// read bearert token from Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			err := echo.ErrUnauthorized
			return session.SetResult(c, nil, stacktrace.CascadeWithClientMessage(err, stacktrace.FORBIDDEN, err.Error()))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			err := echo.ErrUnauthorized
			return session.SetResult(c, nil, stacktrace.CascadeWithClientMessage(err, stacktrace.FORBIDDEN, err.Error()))
		}

		payload, err := utils.ValidateJWT(parts[1], []byte("secret"))
		if err != nil {
			err := echo.ErrUnauthorized
			return session.SetResult(c, nil, stacktrace.CascadeWithClientMessage(err, stacktrace.FORBIDDEN, err.Error()))
		}

		c.Set("user_id", payload.UserID)
		c.Set("name", payload.Name)

		return next(c)
	}
}
