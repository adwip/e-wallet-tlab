package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
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

func (r *WalletHandler) GetBalance(c *echo.Context) error {
	out, err := r.walletsUsecase.GetBalance(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error()))
	}
	return session.SetResult(c, out, nil)
}
