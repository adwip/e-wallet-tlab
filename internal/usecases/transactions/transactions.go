package transactions

import (
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/requests"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions/responses"
	"github.com/labstack/echo/v5"
)

type Transactions interface {
	TopUp(c *echo.Context, req requests.TopUpReq) (out responses.TopUpResp, err error)
	History(c *echo.Context, req requests.TransactionHistoryReq) (out responses.TransactionHistoryListResp, err error)
	Transfer(c *echo.Context, req requests.TransferReq) (out responses.TransferResp, err error)
}
