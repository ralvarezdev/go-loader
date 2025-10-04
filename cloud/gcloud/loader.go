package gcloud

import (
	"context"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"google.golang.org/grpc/credentials/oauth"
)

// LoadGoogleCredentials loads the Google credentials
//
// Parameters:
//
//   - ctx: context.Context
//
// Returns:
//
//   - *google.Credentials: Google credentials
//   - error: error if any
func LoadGoogleCredentials(ctx context.Context) (*google.Credentials, error) {
	credentials, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return nil, ErrFailedToLoadGoogleCredentials
	}
	return credentials, nil
}

// LoadServiceAccountCredentials loads the service account credentials
//
// Parameters:
//
//   - ctx: context.Context
//   - url: string
//   - credentials: *google.Credentials
//
// Returns:
//
//   - *oauth.TokenSource: OAuth2 token source
//   - error: error if any
func LoadServiceAccountCredentials(
	ctx context.Context, url string, credentials *google.Credentials,
) (*oauth.TokenSource, error) {
	// Check if the credentials are nil
	if credentials == nil {
		return nil, ErrNilGoogleCredentials
	}

	// Create a new token source
	tokenSource, err := idtoken.NewTokenSource(
		ctx,
		url,
		option.WithCredentials(credentials),
	)
	if err != nil {
		return nil, ErrFailedToCreateTokenSource
	}

	return &oauth.TokenSource{TokenSource: tokenSource}, nil
}
