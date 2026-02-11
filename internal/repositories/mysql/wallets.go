package mysql

import (
	"github.com/adwip/e-wallet-tlab/internal/models"
	"gorm.io/gorm"
)

type walletRepository struct {
	// put db variables	instances here
	db *gorm.DB
}

func SetupWalletRepository(db *gorm.DB) models.Wallets {
	return &walletRepository{
		db: db,
	}
}
