package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users"
	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	usersUsecase users.UsersUsecase
}

func SetupAuthHandler(usersUsecase users.UsersUsecase) *AuthHandler {
	return &AuthHandler{
		usersUsecase: usersUsecase,
	}
}

func (r *AuthHandler) TestRequest(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}
