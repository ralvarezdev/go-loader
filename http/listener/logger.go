package listener

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the logger for the listener
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger creates a new listener logger
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Check if the logger is nil
	if modeLogger == nil {
		return nil, gologger.ErrNilLogger
	}

	// Initialize the mode named logger
	namedLogger, _ := gologgermodenamed.NewDefaultLogger(header, modeLogger)

	return &Logger{logger: namedLogger}, nil
}

// ServerStarted logs a success message when the server starts
func (l *Logger) ServerStarted(port string) {
	l.logger.Info(
		"server started",
		"port: "+port,
	)
}
