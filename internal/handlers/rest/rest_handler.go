package rest

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session"
	"github.com/labstack/echo/v5"
)

type RestHandler struct {
}

func SetupRestHandlers() *RestHandler {
	return &RestHandler{}
}

func (r *RestHandler) TestRequest(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}
