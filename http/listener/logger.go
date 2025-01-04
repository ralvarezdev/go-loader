package listener

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

// Logger is the logger for the listener
type Logger struct {
	logger gologger.Logger
}

// NewLogger creates a new listener logger
func NewLogger(logger gologger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &Logger{logger: logger}, nil
}

// ServerStarted logs a success message when the server starts
func (l *Logger) ServerStarted(port string) {
	l.logger.LogMessage(
		gologger.NewLogMessage(
			"server started",
			gologgerstatus.Debug,
			port,
		),
	)
}
