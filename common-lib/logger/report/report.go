package report

import (
	"log"
	"sync"
)

// Report is used in usecase level to report error - non blocking flow
type Report interface {
	Info(message string)
	Warning(message string)
	Error(message string)
}

type report struct {
	mu      sync.Mutex
	logging *log.Logger
}

func SetupReport(logger *log.Logger) Report {
	return &report{
		logging: logger,
	}
}
