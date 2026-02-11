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

	receiverWallet, err := r.walletRepo.GetWalletByAccountNumber(req.To)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if receiverWallet.ID == 0 {
		err = errors.New("receiver wallet not found")
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if senderWallet.ID == receiverWallet.ID {
		err = errors.New("sender and receiver wallet cannot be the same")
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if senderWallet.Balance < req.Amount {
		err = errors.New("insufficient balance")
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.INVALID_INPUT, err.Error())
	}

	var operationId = utils.GenerateUUID()
	var transactionIdIN = utils.GenerateUUID()
	var transactionIdOUT = utils.GenerateUUID()
	var status = "FAILED"

	tx := r.db.Begin()

	defer func() {
		if status == "SUCCESS" {
			r.transactionRepo.AddTransactionHistory(entities.TransactionHistories{
				TransactionId: transactionIdIN,
				Amount:        req.Amount,
				SecureId:      utils.GenerateUUID(),
				WalletId:      receiverWallet.SecureId,
				Status:        status,
				Type:          "TRANSFER-IN",
				Description:   req.Note,
			})
		}

		r.transactionRepo.AddTransactionHistory(entities.TransactionHistories{
			TransactionId: transactionIdOUT,
			Amount:        req.Amount,
			SecureId:      utils.GenerateUUID(),
			WalletId:      senderWallet.SecureId,
			Status:        status,
			Type:          "TRANSFER-OUT",
			Description:   req.Note,
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
		WalletID:    senderWallet.SecureId,
		Amount:      req.Amount,
		ActionType:  "TRANSFER-OUT",
		SecureId:    transactionIdOUT,
		OperationId: operationId,
	}

	// act as Ledger
	err = r.transactionRepo.AddTransactionTX(tx, transaction)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	transaction = entities.Transaction{
		WalletID:    receiverWallet.SecureId,
		Amount:      req.Amount,
		ActionType:  "TRANSFER-IN",
		SecureId:    transactionIdIN,
		OperationId: operationId,
	}
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
		SecureId:       utils.GenerateUUID(),
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
