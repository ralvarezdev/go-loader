package env

import (
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the environment logger
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger creates a new environment logger
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Initialize the mode named logger
	namedLogger, err := gologgermodenamed.NewDefaultLogger(header, modeLogger)
	if err != nil {
		return nil, err
	}

	return &Logger{logger: namedLogger}, nil
}

// EnvironmentVariableLoaded logs a message when an environment variable is loaded
func (l *Logger) EnvironmentVariableLoaded(variablesName ...string) {
	l.logger.Debug(
		"environment variable loaded",
		variablesName...,
	)
}
