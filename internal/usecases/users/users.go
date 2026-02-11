package users

import (
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/requests"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/responses"
	"github.com/labstack/echo/v5"
)

type UsersUsecase interface {
	Register(req requests.UserRegistrationReq) (out responses.UserRegistrationResp, err error)
	Login(req requests.UserLoginReq) (out responses.UserLoginResp, err error)
	GetProfile(ctx *echo.Context) (out responses.UserProfileResp, err error)
}
