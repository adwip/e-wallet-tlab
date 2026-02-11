package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions"
	"github.com/adwip/e-wallet-tlab/internal/usecases/wallets"
	"github.com/labstack/echo/v5"
)

type TransactionsHandler struct {
	transactionUsecase transactions.Transactions
	walletUsecase      wallets.WalletsUsecase
}

func SetupTransactionsHandler(transactionUsecase transactions.Transactions, walletUsecase wallets.WalletsUsecase) *TransactionsHandler {
	return &TransactionsHandler{
		transactionUsecase: transactionUsecase,
		walletUsecase:      walletUsecase,
	}
}

func (r *TransactionsHandler) TestRequest(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}

func (r *TransactionsHandler) TopUp(c *echo.Context) error {
	return session.SetResult(c, map[string]interface{}{
		"message": "success",
	}, nil)
}
