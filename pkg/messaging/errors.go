package messaging

import "errors"

var (
	// ErrAlreadyStarted returned when already started.
	ErrAlreadyStarted = errors.New("already started")
	// ErrNotStarted returned when not started.
	ErrNotStarted = errors.New("not started")
	// ErrAlreadyStopped returned when already stopped.
	ErrAlreadyStopped = errors.New("already stopped")
)
