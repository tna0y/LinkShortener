package entities

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotAuthenticated = errors.New("not authenticated")
	ErrNotFound         = errors.New("not found")
	ErrExists           = errors.New("link with given short id exists")
	ErrInvalidShortID   = errors.New("invalid short id")
	ErrInvalidTargetURL = errors.New("invalid target url")
)
