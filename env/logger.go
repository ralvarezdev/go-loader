package env

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

// Logger is the environment logger
type Logger struct {
	logger gologger.Logger
}

// NewLogger creates a new environment logger
func NewLogger(logger gologger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &Logger{logger: logger}, nil
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (l *Logger) EnvironmentVariableLoaded(variablesName ...string) {
	l.logger.LogMessage(
		gologger.NewLogMessage(
			"environment variable loaded",
			gologgerstatus.Debug,
			variablesName...,
		),
	)
}
