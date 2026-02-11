package models

import (
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"gorm.io/gorm"
)

type Wallets interface {
	CreateNewWallet(tx *gorm.DB, wallet entities.Wallet) (err error)
	UpdateBalance(db *gorm.DB, walletId string, amount float64) (err error)
	GetWalletByUserId(userId string) (out entities.Wallet, err error)
	AddNewTransfer(tx *gorm.DB, transfer entities.Transfers) (err error)
	GetWalletByAccountNumber(accountNumber string) (out entities.Wallet, err error)
}
