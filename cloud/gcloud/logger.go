package gcloud

import (
	"log/slog"

	"google.golang.org/grpc/credentials/oauth"
)

// FailedToLoadTokenSource logs the failed to load token source
//
// Parameters:
//
//   - err: the error
//   - logger: the logger
func FailedToLoadTokenSource(err error, logger *slog.Logger) {
	if logger != nil {
		logger.Error(
			"Failed to load token source",
			slog.String("error", err.Error()),
		)
	}
}

// LoadedTokenSource logs the loaded token source
//
// Parameters:
//
//   - tokenSource: the token source
//   - logger: the logger
func LoadedTokenSource(tokenSource *oauth.TokenSource, logger *slog.Logger) {
	// Check if the token source is nil
	if tokenSource == nil {
		FailedToLoadTokenSource(ErrNilTokenSource, logger)
		return
	}

	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		FailedToLoadTokenSource(err, logger)
		return
	}

	if logger != nil {
		logger.Debug(
			"Loaded token source",
			slog.String("access_token", token.AccessToken),
		)
	}
}
