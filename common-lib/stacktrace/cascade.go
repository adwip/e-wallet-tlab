package stacktrace

import (
	"fmt"
	"runtime"
	"strings"
	// _ "github.com/sirupsen/logrus"
)

type stacktraceCode string

type stacktrace struct {
	error
	errFunc   string
	errCode   stacktraceCode
	message   string
	funcStack []string
}

func Cascade(err error, code stacktraceCode, message string) error {
	var (
		stackErr         = make([]uintptr, 10)
		errCast, isStack = err.(stacktrace)
	)
	n := runtime.Callers(2, stackErr)
	frames := runtime.CallersFrames(stackErr[:n])
	frame, _ := frames.Next()

	errFunc := strings.Split(frame.Func.Name(), "/")
	if !isStack {
		errCast = stacktrace{
			error:   err,
			errFunc: fmt.Sprintf("[%s - %d]", errFunc[len(errFunc)-1], frame.Line),
			errCode: code,
		}
		return errCast
	}

	errCast.funcStack = append(errCast.funcStack, fmt.Sprintf("%s - %d", errFunc[len(errFunc)-1], frame.Line))

	return errCast
}

func CascadeWithClientMessage(err error, code stacktraceCode, message string) error {
	var (
		stackErr         = make([]uintptr, 10)
		errCast, isStack = err.(stacktrace)
	)
	n := runtime.Callers(2, stackErr)
	frames := runtime.CallersFrames(stackErr[:n])
	frame, _ := frames.Next()

	errFunc := strings.Split(frame.Func.Name(), "/")
	if !isStack {
		errCast = stacktrace{
			error:   err,
			errFunc: fmt.Sprintf("[%s - %d]", errFunc[len(errFunc)-1], frame.Line),
			errCode: code,
			message: message,
		}
		return errCast
	}

	errCast.funcStack = append(errCast.funcStack, fmt.Sprintf("%s - %d", errFunc[len(errFunc)-1], frame.Line))
	errCast.message = message

	return errCast
}

func DefineStacktrace(err error) (errCode, errFunc, errReason, errMessage string, errStack []string) {
	if err == nil {
		errCode = string(SUCCESS)
		return errCode, errFunc, errReason, errMessage, errStack
	}

	errCast, isStack := err.(stacktrace)
	if !isStack {
		return errCode, errFunc, errReason, errMessage, errStack
	}

	errCode = string(errCast.errCode)
	errFunc = errCast.errFunc
	errReason = errCast.Error()
	errMessage = errCast.message
	errStack = errCast.funcStack

	return errCode, errFunc, errReason, errMessage, errStack
}

func StacktraceToHttpCode(stCode string) (code int, levelCode int, levelName string) {
	code = HttpStatusCodeByStacktrace(stacktraceCode(stCode))
	levelCode, levelName = LevelByStacktrace(stacktraceCode(stCode))
	return code, levelCode, levelName
}

func GetErrorCode(err error) (out string) {
	var (
		errCast = err.(stacktrace)
	)
	return fmt.Sprintf("%s - %s - %s", errCast.errFunc, errCast.errCode, errCast.Error())
}
