package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/requests"
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

func (r *AuthHandler) Login(c *echo.Context) error {
	req, err := requests.NewUserLoginReq(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error()))
	}

	resp, err := r.usersUsecase.Login(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error()))
	}
	return session.SetResult(c, resp, nil)
}
