package rest_session

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/logger"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/shared/constant"
	"github.com/adwip/aj-teknik-backend-admin/common-lib/stacktrace"
)

func (s *restSession) writeRestLog(err error, method, reqId, path string, queryPayload any) (stCode string, stMsg string, errProcess error) {
	var (
		logMsg, errReason   string
		errStack            []string
		logFormat           = "[\x1b[%dm%s\x1b[0m] | %s | %s | %s"
		errMessage, errFunc string
	)
	stCode = string(stacktrace.SUCCESS)
	if err != nil {
		stCode, errFunc, errReason, errMessage, errStack = stacktrace.DefineStacktrace(err)
		if errMessage == "" {
			errMessage = stacktrace.StacktraceMessageByCode(stCode)
		}
	}
	httpCode, colorCode, levelName := stacktrace.StacktraceToHttpCode(stCode)

	logMsg = fmt.Sprintf(logFormat, colorCode, levelName, method, stCode, path)
	// print to bash terminal
	fmt.Println(logMsg)

	// No logging on development environment
	logInfo := logger.LogMap{
		Level:      levelName,
		Method:     method,
		ReqId:      reqId,
		EventTime:  time.Now().UTC().Format(constant.YYMMDDHiS),
		StatusCode: httpCode,
		Path:       path,
		Payload:    queryPayload,
	}

	if err != nil {
		logInfo.ErrorSource = errFunc
		logInfo.ErrorValue = errReason
		logInfo.ErrorStack = errStack
	}

	byteData, errProcess := json.Marshal(logInfo)
	if errProcess != nil {
		return stCode, errMessage, errProcess
	}

	switch levelName {
	case constant.INFO:
		s.log.Info(string(byteData))
	case constant.WARNING:
		s.log.Warning(string(byteData))
	case constant.FATAL:
		s.log.Fatal(string(byteData))
	}

	return stCode, errMessage, nil
}
