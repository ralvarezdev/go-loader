package gcloud

import "errors"

var (
	ErrNilGoogleCredentials          = errors.New("google credentials cannot be nil")
	ErrNilTokenSource                = errors.New("token source cannot be nil")
	ErrFailedToLoadGoogleCredentials = errors.New("failed to load google credentials")
	ErrFailedToCreateTokenSource     = errors.New("failed to create token source")
)
