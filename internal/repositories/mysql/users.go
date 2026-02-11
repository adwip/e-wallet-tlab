package mysql

import (
	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/models"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"gorm.io/gorm"
)

type usersRepository struct {
	// put db variables	instances here
	db *gorm.DB
}

func SetupUsersRepository(db *gorm.DB) models.Users {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) GetUserByEmail(email string) (out entities.Users, err error) {
	if err = r.db.Where("email = ?", email).Scan(&out).Error; err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}

func (r *usersRepository) GetUserBySecureId(secureId string) (out entities.Users, err error) {
	if err = r.db.Where("secure_id = ?", secureId).Scan(&out).Error; err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}

func (r *usersRepository) CreateNewUser(tx *gorm.DB, user entities.Users) (err error) {
	if err = tx.Create(&user).Error; err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}
