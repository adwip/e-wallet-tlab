package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/internal/usecases/wallets"
	"github.com/labstack/echo/v5"
)

type WalletHandler struct {
	walletsUsecase wallets.WalletsUsecase
}

func SetupWalletHandler(walletsUsecase wallets.WalletsUsecase) *WalletHandler {
	return &WalletHandler{
		walletsUsecase: walletsUsecase,
	}
}

func (r *WalletHandler) TestRequest(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}

func (r *WalletHandler) TopUp(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}
