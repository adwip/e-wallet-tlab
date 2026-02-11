package mysql

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/models"
	"github.com/adwip/e-wallet-tlab/internal/models/dtos"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
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

func (r *transactionRepository) AddTransactionTX(tx *gorm.DB, transaction entities.Transaction) (err error) {
	if err = tx.Create(&transaction).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}

func (r *transactionRepository) AddTransaction(transaction entities.Transaction) (err error) {
	if err = r.db.Create(&transaction).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}

func (r *transactionRepository) GetTransactionByID(secureId string) (out entities.Transaction, err error) {
	if err = r.db.Where("secure_id = ?", secureId).First(&out).Error; err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}

func (r *transactionRepository) AddTransactionHistory(transaction entities.TransactionHistories) (err error) {
	if err = r.db.Create(&transaction).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}

func (r *transactionRepository) GetTransactionsByWalletId(walletId string, limit int, offset int) (out []dtos.TransactionHistoryDto, err error) {
	err = r.db.Table("transactions_histories th").
		Select("t.operation_id as operation_id, th.amount, th.status, th.created_at as transaction_date, th.type, th.description").
		Joins("LEFT JOIN transactions t ON t.secure_id = th.transaction_id").
		Where("th.wallet_id = ?", walletId).
		Limit(limit).
		Offset(offset).
		Order("th.created_at desc").
		Find(&out).Error
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}
