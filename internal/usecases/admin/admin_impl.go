package admin

type adminUsecase struct {
}

func SetupAdminUsecase() AdminUsecase {
	return &adminUsecase{}
}
