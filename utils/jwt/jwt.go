package jwt

import (
	"project/consts"
	"project/setting"
	"project/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

// BaseClaims 自定义字段封装进jwt，后续从token想取出什么字段就在这里加
type BaseClaims struct {
	UID      int64
	ID       uint
	Username string
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(setting.Conf.JWT.SigningKey),
	}
}

func (j *JWT) GenerateToken(baseClaims BaseClaims) (accessToken, refreshToken string, err error) {
	jConfig := setting.Conf.JWT
	nowTime := time.Now()
	ae, _ := utils.ParseDuration(jConfig.AccessExpire)
	re, _ := utils.ParseDuration(jConfig.RefreshExpire)
	accessExpireTime := nowTime.Add(ae)
	refreshExpireTime := nowTime.Add(re)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpireTime.Unix(),
			Issuer:    jConfig.Issuer,
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.SigningKey)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: refreshExpireTime.Unix(),
		Issuer:    jConfig.Issuer,
	}).SignedString(j.SigningKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, err
}

// ParseToken 解析token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok {
			return claims, nil
		}
	}

	return nil, err
}

// ParseRefreshToken 以旧换新
func (j *JWT) ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	accessClaims, err := j.ParseToken(aToken)
	if err != nil {
		return
	}

	refreshClaim, err := j.ParseToken(rToken)
	if err != nil {
		return
	}
	aClaim := BaseClaims{
		UID:      accessClaims.UID,
		ID:       accessClaims.ID,
		Username: accessClaims.Username,
	}
	if accessClaims.ExpiresAt > time.Now().Unix() {

		return j.GenerateToken(aClaim)
	}

	if refreshClaim.ExpiresAt > time.Now().Unix() {

		return j.GenerateToken(aClaim)
	}

	return "", "", consts.TokenExpired
}
