package routes

import (
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/web"
	"github.com/labstack/echo/v5"
)

type webRoutes struct {
	web *web.WebHandler
}

func SetupWebRoutes(web *web.WebHandler) *webRoutes {
	return &webRoutes{
		web: web,
	}
}

func (w *webRoutes) RegisterRoutes(g *echo.Group) {

}
