package consts

import (
	"errors"

	"golang.org/x/sync/singleflight"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
	TokenNotValidYet = errors.New("token not active yet")
	SingleFlight     = &singleflight.Group{}
)

const (
	AccessTokenHeader    = "access_token"
	RefreshTokenHeader   = "refresh_token"
	HeaderForwardedProto = "X-Forwarded-Proto"
	MaxAge               = 3600 * 24
)
