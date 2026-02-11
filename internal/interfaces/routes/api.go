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
	g.POST("/auth/register", a.usersHandlers.Register)
	g.POST("/auth/login", a.authHandlers.Login)
}

func (a *apiRoutes) ProtectedRoutes(g *echo.Group) {
	g.Use(middlewares.LoginValidation)
	g.GET("/users/profile", a.usersHandlers.GetProfile)
	g.POST("/wallets/topup", a.transactionHandlers.TopUp)
	g.GET("/wallets/balance", a.walletHandlers.GetBalance)
	g.POST("/transactions/transfer", a.transactionHandlers.Transfer)
	g.GET("/transactions/history", a.transactionHandlers.History)
}
