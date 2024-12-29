package listener

import "errors"

var (
	ErrFailedToClose  = errors.New("failed to close listener")
	ErrFailedToListen = errors.New("failed to listen")
	ErrFailedToServe  = errors.New("failed to serve")
)
