package users

import (
	"errors"

	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/adwip/e-wallet-tlab/internal/shared/utils"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/requests"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users/responses"
)

func (s *usersUsecase) Login(req requests.UserLoginReq) (out responses.UserLoginResp, err error) {

	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if user.ID == 0 {
		err = errors.New("Email not found")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if err = utils.VerifyPassword(req.Password, user.Password); err != nil {
		return out, stacktrace.Cascade(err, stacktrace.UNAUTHENTICATED, err.Error())
	}
	jwtPayload := utils.PayloadSchema{
		UserID: user.SecureId,
		Name:   user.Name,
	}
	token, err := utils.GenerateJWT(jwtPayload, []byte("secret"))
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = responses.UserLoginResp{
		UserId: user.SecureId,
		Name:   user.Name,
		Token:  token,
	}
	return out, nil
}
