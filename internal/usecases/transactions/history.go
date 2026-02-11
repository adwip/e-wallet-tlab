package transactions

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/requests"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/responses"
	"github.com/labstack/echo/v5"
)

func (r *transactionUsecase) History(c *echo.Context, req requests.TransactionHistoryReq) (out responses.TransactionHistoryListResp, err error) {
	wallet, err := r.walletRepo.GetWalletByUserId(utils.GetUserId(c))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	transactions, err := r.transactionRepo.GetTransactionsByWalletId(wallet.SecureId, req.Limit, req.Offset)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = responses.TransactionHistoryListResp{
		WalletId: wallet.SecureId,
	}

	for _, transaction := range transactions {
		out.Transactions = append(out.Transactions, responses.TransactionHistoryResp{
			TransactionId:   transaction.TransactionId,
			Amount:          transaction.Amount,
			Status:          transaction.Status,
			TransactionDate: transaction.TransactionDate.Format("2006-01-02 15:04:05"),
			Type:            transaction.Type,
			Description:     transaction.Description,
		})
	}
	return out, nil
}
