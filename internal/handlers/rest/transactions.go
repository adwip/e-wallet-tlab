package rest

import (
	"github.com/adwip/e-wallet-tlab/common-lib/session"
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/requests"
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

func (r *TransactionsHandler) TopUp(c *echo.Context) error {
	req, err := requests.NewTopUpReq(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error()))
	}

	out, err := r.transactionUsecase.TopUp(c, req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error()))
	}
	return session.SetResult(c, out, nil)
}

func (r *TransactionsHandler) Transfer(c *echo.Context) error {
	req, err := requests.NewTransferReq(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error()))
	}

	out, err := r.transactionUsecase.Transfer(c, req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error()))
	}
	return session.SetResult(c, out, nil)
}

func (r *TransactionsHandler) History(c *echo.Context) error {
	req, err := requests.NewTransactionHistoryReq(c)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error()))
	}

	out, err := r.transactionUsecase.History(c, req)
	if err != nil {
		return session.SetResult(c, nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error()))
	}
	return session.SetResult(c, out, nil)
}
