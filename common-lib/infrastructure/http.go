package infrastructure

import (
	"fmt"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/logger"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/session/rest_session"
	"github.com/labstack/echo/v5"
)

type HttpServer interface {
	StartServer(port string) error
	RouteInit(prefix string) *echo.Group
}

type httpServer struct {
	server *echo.Echo
}

func SetupHttpServer(log logger.Logger) HttpServer {

	server := echo.New()

	restSession := rest_session.SetupRestSession(log)
	server.Use(restSession.ResultInterceptor)
	server.HTTPErrorHandler = restSession.ErrorHandler
	return &httpServer{
		server: server,
	}
}

func (r *httpServer) RouteInit(prefix string) *echo.Group {
	if prefix == "" {
		prefix = "/api"
	}
	return r.server.Group(prefix)
}

func (r *httpServer) StartServer(port string) error {
	for _, route := range r.server.Router().Routes() {
		fmt.Printf("[%s] - %s - %s \n", route.Method, route.Path, route.Name)
	}
	return r.server.Start(port)
}
