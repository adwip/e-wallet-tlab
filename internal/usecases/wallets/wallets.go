package wallets

import (
	"github.com/adwip/e-wallet-tlab/internal/usecases/wallets/responses"
	"github.com/labstack/echo/v5"
)

type WalletsUsecase interface {
	GetBalance(c *echo.Context) (out responses.GetBalanceResp, err error)
}
