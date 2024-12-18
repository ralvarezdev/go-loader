package gcloud

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	"google.golang.org/grpc/credentials/oauth"
)

// Logger is the logger for Google Cloud
type Logger struct {
	logger gologger.Logger
}

// NewLogger is the logger for Google Cloud
func NewLogger(logger gologger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// LoadedTokenSource logs the loaded token source
func (l *Logger) LoadedTokenSource(tokenSource *oauth.TokenSource) {
	// Check if the token source is nil
	if tokenSource == nil {
		l.logger.LogError(gologger.NewLogError("failed to load token source", nil, NilTokenSourceError))
		return
	}

	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		l.logger.LogError(gologger.NewLogError("failed to load token source", nil, err))
		return
	}

	l.logger.LogMessage(
		gologger.NewLogMessage(
			"loaded token source",
			gologgerstatus.StatusDebug,
			nil,
			token.AccessToken,
		),
	)
}
