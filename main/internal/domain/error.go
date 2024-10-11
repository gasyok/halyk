package domain

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrServerInternal  = errors.New("internal error")
	ErrInsufficientLot = errors.New("insufficient lot")
)
