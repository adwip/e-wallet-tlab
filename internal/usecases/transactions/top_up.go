package transactions

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/requests"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/responses"
	"github.com/labstack/echo/v5"
)

func (r *transactionUsecase) TopUp(c *echo.Context, req requests.TopUpReq) (out responses.TopUpResp, err error) {

	wallet, err := r.walletRepo.GetWalletByUserId(utils.GetUserId(c))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	var amount = wallet.Balance + req.Amount
	var operationId = utils.GenerateUUID()
	var transactionID = utils.GenerateUUID()
	var status = "FAILED"

	tx := r.db.Begin()

	defer func() {
		r.transactionRepo.AddTransactionHistory(entities.TransactionHistories{
			TransactionId: transactionID,
			WalletId:      wallet.SecureId,
			Amount:        req.Amount,
			SecureId:      utils.GenerateUUID(),
			Status:        status,
			Type:          "TOP_UP",
			Description:   req.Note,
		})

		if status != "SUCCESS" {
			amount = wallet.Balance
		}

		out = responses.TopUpResp{
			TransactionId: operationId,
			Balance:       amount,
			Status:        status,
		}
	}()

	err = r.walletRepo.UpdateBalance(tx, wallet.SecureId, amount)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	transaction := entities.Transaction{
		WalletID:    wallet.SecureId,
		SecureId:    transactionID,
		Amount:      req.Amount,
		ActionType:  "TOP_UP",
		OperationId: operationId,
		Note:        req.Note,
	}

	// act as Ledger
	err = r.transactionRepo.AddTransactionTX(tx, transaction)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	wallet, err = r.walletRepo.GetWalletByUserId(utils.GetUserId(c))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	status = "SUCCESS"
	return out, nil
}
