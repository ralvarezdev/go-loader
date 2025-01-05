package gcloud

import (
	gologgermode "github.com/ralvarezdev/go-logger/mode"
	gologgermodenamed "github.com/ralvarezdev/go-logger/mode/named"
	"google.golang.org/grpc/credentials/oauth"
)

// Logger is the logger for Google Cloud
type Logger struct {
	logger gologgermodenamed.Logger
}

// NewLogger is the logger for Google Cloud
func NewLogger(header string, modeLogger gologgermode.Logger) (*Logger, error) {
	// Initialize the mode named logger
	namedLogger, err := gologgermodenamed.NewDefaultLogger(header, modeLogger)
	if err != nil {
		return nil, err
	}

	return &Logger{logger: namedLogger}, nil
}

// FailedToLoadTokenSource logs the failed to load token source
func (l *Logger) FailedToLoadTokenSource(err error) {
	l.logger.Error(
		"failed to load token source",
		err,
	)
}

// LoadedTokenSource logs the loaded token source
func (l *Logger) LoadedTokenSource(tokenSource *oauth.TokenSource) {
	// Check if the token source is nil
	if tokenSource == nil {
		l.FailedToLoadTokenSource(ErrNilTokenSource)
		return
	}

	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		l.FailedToLoadTokenSource(err)
		return
	}

	l.logger.Debug(
		"loaded token source",
		token.AccessToken,
	)
}
