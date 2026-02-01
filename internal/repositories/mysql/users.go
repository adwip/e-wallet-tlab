package mysql

import "github.com/adwip/aj-teknik-backend-admin/internal/model"

type usersRepository struct {
	// put db variables	instances here
}

func SetupUsersRepository() model.Users {
	return &usersRepository{}
}
