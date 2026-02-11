package rest_session

import (
	"fmt"

	"github.com/adwip/e-wallet-tlab/common-lib/logger"
	"github.com/adwip/e-wallet-tlab/common-lib/metadata"
	"github.com/labstack/echo/v5"
)

type RestSession interface {
	ResultInterceptor(next echo.HandlerFunc) echo.HandlerFunc
	ErrorHandler(ctx *echo.Context, err error)
}

type restSession struct {
	log logger.Logger
}

func SetupRestSession(log logger.Logger) RestSession {
	return &restSession{
		log: log,
	}
}

func (s *restSession) ResultInterceptor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx *echo.Context) (err error) {
		// calling the next handler
		err = next(ctx)

		_, _, errExtraction := s.writeRestLog(err, ctx.Request().Method, ctx.Request().Header.Get(metadata.XRequestId), ctx.Path())
		if errExtraction != nil {
			s.log.Fatal(fmt.Sprintf("[\x1b[%dm%s\x1b[0m] %s \n", 31, "FATAL", errExtraction.Error()))
		}
		return nil
	}
}
