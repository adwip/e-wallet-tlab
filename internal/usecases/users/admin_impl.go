package users

import "github.com/adwip/e-wallet-tlab/internal/models"

type usersUsecase struct {
	userRepo models.Users
}

func SetupUsersUsecase(userRepo models.Users) UsersUsecase {
	return &usersUsecase{
		userRepo: userRepo,
	}
}
