package transactions

import "github.com/adwip/e-wallet-tlab/internal/models"

type transactionUsecase struct {
	// put db variables	instances here
	transactionRepo models.Transactions
	walletRepo      models.Wallets
}

func SetupTransactionUsecase(transactionRepo models.Transactions, walletRepo models.Wallets) Transactions {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		walletRepo:      walletRepo,
	}
}
