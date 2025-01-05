package env

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the environment logger
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger creates a new environment logger
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Check if the logger is nil
	if modeLogger == nil {
		return nil, gologger.ErrNilLogger
	}

	// Initialize the mode named logger
	namedLogger, _ := gologgermodenamed.NewDefaultLogger(header, modeLogger)

	return &Logger{logger: namedLogger}, nil
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (l *Logger) EnvironmentVariableLoaded(variablesName ...string) {
	l.logger.Debug(
		"environment variable loaded",
		variablesName...,
	)
}
