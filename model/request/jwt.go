package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

// BaseClaims 自定义字段封装进jwt，后续从token想取出什么字段就在这里加
type BaseClaims struct {
	UID      int64
	ID       uint
	Username string
}
