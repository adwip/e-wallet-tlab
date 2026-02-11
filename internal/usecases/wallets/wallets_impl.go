package wallets

import "github.com/adwip/e-wallet-tlab/internal/models"

type walletsUsecase struct {
	walletRepo models.Wallets
}

func SetupWalletsUsecase(walletRepo models.Wallets) WalletsUsecase {
	return &walletsUsecase{
		walletRepo: walletRepo,
	}
}
