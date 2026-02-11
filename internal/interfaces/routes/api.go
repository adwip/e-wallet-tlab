package routes

import (
	"github.com/adwip/e-wallet-tlab/internal/handlers/rest"
	"github.com/adwip/e-wallet-tlab/internal/interfaces/middlewares"
	"github.com/labstack/echo/v5"
)

type apiRoutes struct {
	authHandlers        *rest.AuthHandler
	usersHandlers       *rest.UsersHandler
	walletHandlers      *rest.WalletHandler
	transactionHandlers *rest.TransactionsHandler
}

func SetupApiRoutes(authHandlers *rest.AuthHandler, usersHandlers *rest.UsersHandler, walletHandlers *rest.WalletHandler, transactionHandlers *rest.TransactionsHandler) *apiRoutes {
	return &apiRoutes{
		authHandlers:        authHandlers,
		usersHandlers:       usersHandlers,
		walletHandlers:      walletHandlers,
		transactionHandlers: transactionHandlers,
	}
}

func (a *apiRoutes) UnProtectedRoutes(g *echo.Group) {
	g.POST("/auth/register", a.authHandlers.TestRequest)
	g.POST("/auth/login", a.authHandlers.TestRequest)
}

func (a *apiRoutes) ProtectedRoutes(g *echo.Group) {
	g.Use(middlewares.LoginValidation)
	g.GET("/users/profile", a.usersHandlers.TestRequest)
	g.POST("/wallets/topup", a.walletHandlers.TopUp)
	g.GET("/wallets/balance", a.walletHandlers.TestRequest)
	g.POST("/transactions/transfer", a.transactionHandlers.TestRequest)
	g.GET("/transactions/history", a.transactionHandlers.TestRequest)
}
