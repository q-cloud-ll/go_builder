package middlewares

import (
	"project/consts"
	"project/utils/app"
	"project/utils/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("access_token")
		refreshToken := c.GetHeader("refresh_token")
		if accessToken == "" {
			app.ResponseErrorWithCodeMsg(c, app.CodeNoLogin, "未登陆")
			c.Abort()
			return
		}
		j := jwt.NewJWT()
		newAccessToken, newRefreshToken, err := j.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			app.ResponseErrorWithCodeMsg(c, app.CodeLoginExpire, err.Error())
			c.Abort()
			return
		}
		claims, err := j.ParseToken(newAccessToken)
		if err != nil {
			app.ResponseErrorWithCodeMsg(c, app.CodeErrorAuthCheckTokenFail, err.Error())
			c.Abort()
			return
		}
		SetToken(c, newAccessToken, newRefreshToken)
		c.Request = c.Request.WithContext(app.NewContext(c.Request.Context(), &app.UserInfo{Username: claims.Username, ID: claims.ID, UID: claims.UID}))

		c.Next()
	}
}

func SetToken(c *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(c)
	c.Header(consts.AccessTokenHeader, accessToken)
	c.Header(consts.RefreshTokenHeader, refreshToken)
	c.SetCookie(consts.AccessTokenHeader, accessToken, consts.MaxAge, "/", "", secure, true)
	c.SetCookie(consts.RefreshTokenHeader, refreshToken, consts.MaxAge, "/", "", secure, true)
}

func IsHttps(c *gin.Context) bool {
	if c.GetHeader(consts.HeaderForwardedProto) == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}
