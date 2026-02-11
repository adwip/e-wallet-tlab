package containers

import (
	"github.com/adwip/e-wallet-tlab/common-lib/infrastructure"
	"github.com/adwip/e-wallet-tlab/common-lib/logger"
	"github.com/adwip/e-wallet-tlab/internal/handlers/rest"
	"github.com/adwip/e-wallet-tlab/internal/interfaces/drivers"
	"github.com/adwip/e-wallet-tlab/internal/interfaces/routes"
	"github.com/adwip/e-wallet-tlab/internal/repositories/mysql"
	"github.com/adwip/e-wallet-tlab/internal/shared/config"
	"github.com/adwip/e-wallet-tlab/internal/usecases/transactions"
	"github.com/adwip/e-wallet-tlab/internal/usecases/users"
	"github.com/adwip/e-wallet-tlab/internal/usecases/wallets"
)

func SetupServiceContainer() (err error) {

	cfg, err := config.SetupConfig()
	if err != nil {
		return nil
	}

	log, _, err := logger.SetupLogger(cfg.Service.LogFile)
	if err != nil {
		return nil
	}

	httpServer := infrastructure.SetupHttpServer(log)

	db, err := drivers.SetupDatabase(cfg.Db.Host)
	if err != nil {
		return nil
	}

	// setup repo
	usersRepo := mysql.SetupUsersRepository(db)
	transactionRepo := mysql.SetupTransactionRepository(db)
	walletRepo := mysql.SetupWalletRepository(db)

	// setup usecase
	usersUsecase := users.SetupUsersUsecase(usersRepo)
	transactionUsecase := transactions.SetupTransactionUsecase(transactionRepo, walletRepo)
	walletUsecase := wallets.SetupWalletsUsecase(walletRepo)

	// setup handler
	authHandler := rest.SetupAuthHandler(usersUsecase)
	usersHandlers := rest.SetupUsersHandler(usersUsecase)
	walletHandlers := rest.SetupWalletHandler(walletUsecase)
	transactionHandlers := rest.SetupTransactionsHandler(transactionUsecase, walletUsecase)

	routes.SetupRoutes(authHandler, usersHandlers, walletHandlers, transactionHandlers, httpServer)

	err = httpServer.StartServer(cfg.Service.Port)
	if err != nil {
		return err
	}
	return nil
}
