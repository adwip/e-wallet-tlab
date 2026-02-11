package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users"
	"github.com/labstack/echo/v5"
)

type UsersHandler struct {
	usersUsecase users.UsersUsecase
}

func SetupUsersHandler(usersUsecase users.UsersUsecase) *UsersHandler {
	return &UsersHandler{
		usersUsecase: usersUsecase,
	}
}

func (r *UsersHandler) TestRequest(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}
