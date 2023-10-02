package util

import (
	"errors"
)

var (
	ErrInvalidJWT = errors.New("invalid jwt")
	ErrInvalidID  = errors.New("invalid id")
)
