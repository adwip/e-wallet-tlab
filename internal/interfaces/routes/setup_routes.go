package routes

import (
	"github.com/adwip/e-wallet-tlab/common-lib/infrastructure"
	"github.com/adwip/e-wallet-tlab/internal/handlers/rest"
)

func SetupRoutes(authHandlers *rest.AuthHandler, usersHandlers *rest.UsersHandler, walletHandlers *rest.WalletHandler, transactionHandlers *rest.TransactionsHandler, server infrastructure.HttpServer) {
	apiRoutes := SetupApiRoutes(authHandlers, usersHandlers, walletHandlers, transactionHandlers)
	api := server.RouteInit("/api")
	apiRoutes.UnProtectedRoutes(api)
	apiRoutes.ProtectedRoutes(api)
}
