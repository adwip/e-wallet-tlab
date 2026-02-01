package rest_session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/logger"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/metadata"
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
		buf1, errExtraction := io.ReadAll(ctx.Request().Body)
		if errExtraction != nil {
			s.log.Fatal(fmt.Sprintf("[\x1b[%dm%s\x1b[0m] %s \n", 31, "FATAL", errExtraction.Error()))
		}

		ctx.Request().Body = io.NopCloser(bytes.NewBuffer(buf1))

		// calling the next handler
		err = next(ctx)

		byteData, errExtraction := json.Marshal(ctx.QueryParams())
		if errExtraction != nil {
			s.log.Fatal(fmt.Sprintf("[\x1b[%dm%s\x1b[0m] %s \n", 31, "FATAL", errExtraction.Error()))
		}
		var payload = string(byteData)
		if ctx.Request().Method != http.MethodGet {
			body, errExtraction := io.ReadAll(io.NopCloser(bytes.NewBuffer(buf1)))
			if errExtraction != nil {
				s.log.Fatal(fmt.Sprintf("[\x1b[%dm%s\x1b[0m] %s \n", 31, "FATAL", errExtraction.Error()))
			}
			payload = string(body)
		}

		_, _, errExtraction = s.writeRestLog(err, ctx.Request().Method, ctx.Request().Header.Get(metadata.XRequestId), ctx.Path(), payload)
		if errExtraction != nil {
			s.log.Fatal(fmt.Sprintf("[\x1b[%dm%s\x1b[0m] %s \n", 31, "FATAL", errExtraction.Error()))
		}
		return nil
	}
}
