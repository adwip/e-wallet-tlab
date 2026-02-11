package users

import (
	"github.com/adwip/e-wallet-tlab/internal/models"
	"gorm.io/gorm"
)

type usersUsecase struct {
	userRepo   models.Users
	db         *gorm.DB
	walletRepo models.Wallets
}

func SetupUsersUsecase(userRepo models.Users, db *gorm.DB, walletRepo models.Wallets) UsersUsecase {
	return &usersUsecase{
		userRepo:   userRepo,
		db:         db,
		walletRepo: walletRepo,
	}
}
