package wallets

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/adwip/e-wallet-tlab/internal/usecases/wallets/responses"
	"github.com/labstack/echo/v5"
)

func (r *walletsUsecase) GetBalance(c *echo.Context) (out responses.GetBalanceResp, err error) {
	wallet, err := r.walletRepo.GetWalletByUserId(utils.GetUserId(c))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = responses.GetBalanceResp{
		Balance: wallet.Balance,
	}
	return out, nil
}
