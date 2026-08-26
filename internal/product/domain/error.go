package domain

import "errors"

var (
	ErrProductCannotBeCancelled = errors.New(
		"product cannot be cancelled",
	)
)
