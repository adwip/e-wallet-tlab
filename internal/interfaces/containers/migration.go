package containers

import (
	"github.com/adwip/e-wallet-tlab/internal/interfaces/drivers"
	"github.com/adwip/e-wallet-tlab/internal/models/entities"
	"github.com/adwip/e-wallet-tlab/internal/shared/config"
)

func Migrations() (err error) {
	cfg, err := config.SetupConfig()
	if err != nil {
		return err
	}

	db, err := drivers.SetupDatabase(cfg.Db.Host)
	if err != nil {
		return err
	}

	// Auto migrate
	err = db.AutoMigrate(&entities.Users{}, &entities.Wallet{}, &entities.Transaction{})
	if err != nil {
		return err
	}

	return nil
}
