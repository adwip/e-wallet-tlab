package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/requests"
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

func (r *UsersHandler) Register(c *echo.Context) error {
	req, err := requests.NewUserRegistrationReq(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error()))
	}

	out, err := r.usersUsecase.Register(req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error()))
	}

	return session.SetResult(c, out, nil)
}

func (r *UsersHandler) GetProfile(c *echo.Context) error {
	out, err := r.usersUsecase.GetProfile(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error()))
	}
	return session.SetResult(c, out, nil)
}

func (r *UsersHandler) TestRequest(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}
