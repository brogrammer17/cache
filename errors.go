package cache

import "errors"

var (
	ErrNotFound   = errors.New("key not found")
	ErrExpired    = errors.New("key expired")
	ErrInvalidTTL = errors.New("ttl must be >= 0")
)
