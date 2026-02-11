package transactions

import (
	"github.com/adwip/e-wallet-tlab/internal/models"
	"gorm.io/gorm"
)

type transactionUsecase struct {
	// put db variables	instances here
	transactionRepo models.Transactions
	walletRepo      models.Wallets
	db              *gorm.DB
}

func SetupTransactionUsecase(transactionRepo models.Transactions, walletRepo models.Wallets, db *gorm.DB) Transactions {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		walletRepo:      walletRepo,
		db:              db,
	}
}
