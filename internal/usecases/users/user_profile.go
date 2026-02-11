package users

import (
	"errors"

	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/responses"
	"github.com/labstack/echo/v5"
)

func (s *usersUsecase) GetProfile(ctx *echo.Context) (out responses.UserProfileResp, err error) {
	user, err := s.userRepo.GetUserBySecureId(utils.GetUserId(ctx))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if user.ID == 0 {
		err = errors.New("User not found")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	wallet, err := s.walletRepo.GetWalletByUserId(utils.GetUserId(ctx))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = responses.UserProfileResp{
		UserID:        user.SecureId,
		Name:          user.Name,
		Email:         user.Email,
		AccountNumber: wallet.AccountNumber,
	}
	return out, nil
}
