package models

import (
	"github.com/adwip/e-wallet-tlab/internal/models/dtos"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"gorm.io/gorm"
)

type Transactions interface {
	AddTransactionTX(tx *gorm.DB, transaction entities.Transaction) (err error)
	AddTransaction(transaction entities.Transaction) (err error)
	GetTransactionByID(secureId string) (out entities.Transaction, err error)
	AddTransactionHistory(transaction entities.TransactionHistories) (err error)
	GetTransactionsByWalletId(walletId string, limit int, offset int) (out []dtos.TransactionHistoryDto, err error)
}
