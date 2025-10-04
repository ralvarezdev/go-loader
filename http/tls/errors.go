package tls

import (
	"errors"
)

var (
	ErrFailedToAddCAPem              = errors.New("failed to add server ca's certificate")
	ErrFailedToLoadSystemCredentials = errors.New("failed to load system credentials")
)
