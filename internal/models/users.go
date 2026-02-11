package models

import (
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"gorm.io/gorm"
)

type Users interface {
	GetUserByEmail(email string) (out entities.Users, err error)
	CreateNewUser(tx *gorm.DB, user entities.Users) (err error)
	GetUserBySecureId(secureId string) (out entities.Users, err error)
}
