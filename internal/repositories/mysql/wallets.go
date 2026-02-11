package mysql

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/models"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *walletRepository) CreateNewWallet(tx *gorm.DB, wallet entities.Wallet) (err error) {
	if err = tx.Create(&wallet).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}

func (r *walletRepository) GetWalletByUserId(userId string) (out entities.Wallet, err error) {
	if err = r.db.Where("user_id = ?", userId).First(&out).Error; err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}

func (r *walletRepository) UpdateBalance(db *gorm.DB, walletId string, amount float64) (err error) {

	var wallet entities.Wallet

	if err := db.Clauses(clause.Locking{Strength: "UPDATE"}).First(&wallet, "secure_id = ?", walletId).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	wallet.Balance = amount

	if err := db.Save(&wallet).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	return nil
}

func (r *walletRepository) AddNewTransfer(tx *gorm.DB, transfer entities.Transfers) (err error) {
	if err = tx.Create(&transfer).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}
