package users

import (
	"errors"

	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/requests"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/responses"
)

func (s *usersUsecase) Register(req requests.UserRegistrationReq) (out responses.UserRegistrationResp, err error) {
	existing, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if existing.ID != 0 {
		err = errors.New("Email already used")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	user := entities.Users{
		SecureId: utils.GenerateUUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	wallet := entities.Wallet{
		SecureId:      utils.GenerateUUID(),
		UserID:        user.SecureId,
		Balance:       0,
		AccountNumber: utils.GenerateAccountNumber(),
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return out, stacktrace.Cascade(tx.Error, stacktrace.INTERNAL_SERVER_ERROR, tx.Error.Error())
	}

	err = s.userRepo.CreateNewUser(tx, user)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	err = s.walletRepo.CreateNewWallet(tx, wallet)
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = responses.UserRegistrationResp{
		SecureId: user.SecureId,
		Name:     user.Name,
	}
	return out, nil
}
