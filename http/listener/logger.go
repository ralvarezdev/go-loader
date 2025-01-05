package listener

import (
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
)

// Logger is the logger for the listener
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger creates a new listener logger
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Initialize the mode named logger
	namedLogger, err := gologgermodenamed.NewDefaultLogger(header, modeLogger)
	if err != nil {
		return nil, err
	}

	return &Logger{logger: namedLogger}, nil
}

// ServerStarted logs a success message when the server starts
func (l *Logger) ServerStarted(port string) {
	l.logger.Info(
		"server started",
		"port: "+port,
	)
}
