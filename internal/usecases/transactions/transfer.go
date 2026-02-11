package transactions

import (
	"errors"

	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/requests"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/responses"
	"github.com/labstack/echo/v5"
)

func (r *transactionUsecase) Transfer(c *echo.Context, req requests.TransferReq) (out responses.TransferResp, err error) {
	senderWallet, err := r.walletRepo.GetWalletByUserId(utils.GetUserId(c))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	receiverWallet, err := r.walletRepo.GetWalletByUserId(utils.GetUserId(c))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if senderWallet.ID == receiverWallet.ID {
		err = errors.New("sender and receiver wallet cannot be the same")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if senderWallet.Balance < req.Amount {
		err = errors.New("insufficient balance")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	var operationId = utils.GenerateUUID()
	var status = "FAILED"

	tx := r.db.Begin()

	defer func() {
		r.transactionRepo.AddTransactionHistory(entities.TransactionHistories{
			TransactionId: operationId,
			Status:        status,
		})

		out = responses.TransferResp{
			TransactionId: operationId,
			Status:        status,
		}
	}()

	err = r.walletRepo.UpdateBalance(tx, senderWallet.SecureId, senderWallet.Balance-req.Amount)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	transaction := entities.Transaction{
		WalletID:   senderWallet.ID,
		Amount:     req.Amount,
		ActionType: "TRANSFER",
		SecureId:   operationId,
	}

	// act as Ledger
	err = r.transactionRepo.AddTransactionTX(tx, transaction)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	err = r.walletRepo.UpdateBalance(tx, receiverWallet.SecureId, receiverWallet.Balance+req.Amount)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	transfer := entities.Transfers{
		SenderId:       senderWallet.UserID,
		WalletSourceId: senderWallet.SecureId,
		ReceiverId:     receiverWallet.UserID,
		WalletDestId:   receiverWallet.SecureId,
		Amount:         req.Amount,
		SecureId:       operationId,
		Note:           req.Note,
	}
	err = r.walletRepo.AddNewTransfer(tx, transfer)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	status = "SUCCESS"
	return out, nil
}
