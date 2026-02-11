package mysql

import (
	"github.com/adwip/e-wallet-tlab/internal/models"
	"gorm.io/gorm"
)

type transactionRepository struct {
	// put db variables	instances here
	db *gorm.DB
}

func SetupTransactionRepository(db *gorm.DB) models.Transactions {
	return &transactionRepository{
		db: db,
	}
}
