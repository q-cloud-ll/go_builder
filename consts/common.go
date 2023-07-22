package consts

import (
	"errors"

	"golang.org/x/sync/singleflight"
)

const (
	ConfigEnv         = "GB_CONFIG"
	ConfigDefaultFile = "conf/config.yaml"
	ConfigTestFile    = "conf/config.test.yaml"
	ConfigReleaseFile = "conf/config.release.yaml"
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

const (
	UploadMaxM           = 10
	UploadMaxBytes int64 = 1024 * 1024 * 1024 * UploadMaxM
)

const SpanCTX = "span-ctx"

var (
	GetTrackIdErr = errors.New("获取 track id 错误")
)

const (
	OtherSen = "./static/dictionary/其他词库.txt"
)
