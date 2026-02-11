package wallets

import (
	"github.com/adwip/e-wallet-tlab/internal/models"
	"gorm.io/gorm"
)

type walletsUsecase struct {
	walletRepo      models.Wallets
	transactionRepo models.Transactions
	db              *gorm.DB
}

func SetupWalletsUsecase(walletRepo models.Wallets, transactionRepo models.Transactions, db *gorm.DB) WalletsUsecase {
	return &walletsUsecase{
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
		db:              db,
	}
}
