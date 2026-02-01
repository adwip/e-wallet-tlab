package logger

import (
	"log"
	"os"
	"sync"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/logger/report"
)

type Logger interface {
	Info(message string)
	Warning(message string)
	Fatal(message string)
}

type logger struct {
	mu      sync.Mutex
	logging *log.Logger
}

func SetupLogger(logPath string) (Logger, report.Report, error) {
	var (
		logFile = os.Stdout
		err     error
	)

	if logPath != "" {
		logFile, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, nil, err
		}
	}

	logging := log.New(logFile, "\n", log.Lmsgprefix|log.Ldate|log.Ltime|log.LUTC)

	report := report.SetupReport(logging)

	return &logger{
		logging: logging,
	}, report, nil
}

func (l *logger) Info(message string) {
	l.writeLog(message)
}

func (l *logger) Warning(message string) {
	l.writeLog(message)
}

func (l *logger) Fatal(message string) {
	l.writeLog(message)
}
