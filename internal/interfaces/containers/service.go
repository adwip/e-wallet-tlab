package containers

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/infrastructure"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/logger"
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/rest"
	"github.com/adwip/aj-teknik-backend-admin/internal/handlers/web"
	"github.com/adwip/aj-teknik-backend-admin/internal/interfaces/drivers"
	"github.com/adwip/aj-teknik-backend-admin/internal/interfaces/routes"
	"github.com/adwip/aj-teknik-backend-admin/internal/repositories/mysql"
	"github.com/adwip/aj-teknik-backend-admin/internal/shared/config"
	"github.com/adwip/aj-teknik-backend-admin/internal/usecases/admin"
)

func SetupServiceContainer() (err error) {
	log, _, err := logger.SetupLogger("")
	if err != nil {
		return nil
	}

	httpServer := infrastructure.SetupHttpServer(log)

	_, err = config.SetupConfig()
	if err != nil {
		return nil
	}

	err = drivers.SetupDatabase("config.DatabaseUrl")
	if err != nil {
		return nil
	}

	// setup repo
	_ = mysql.SetupUsersRepository()

	// setup usecase
	_ = admin.SetupAdminUsecase()

	// setup handler
	rest := rest.SetupRestHandlers()
	web := web.SetupWebHandlers()

	routes.SetupRoutes(rest, web, httpServer)

	err = httpServer.StartServer(":8080")
	if err != nil {
		return err
	}
	return nil
}
