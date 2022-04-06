package jwt

import "errors"

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("malformed token")
	ErrTokenInvalid     = errors.New("invalid token")
)
