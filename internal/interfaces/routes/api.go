package routes

import (
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/rest"
	"github.com/labstack/echo/v5"
)

type apiRoutes struct {
	rest *rest.RestHandler
}

func SetupApiRoutes(rest *rest.RestHandler) *apiRoutes {
	return &apiRoutes{
		rest: rest,
	}
}

func (a *apiRoutes) RegisterRoutes(g *echo.Group) {
	g.GET("/test", a.rest.TestRequest)
}
