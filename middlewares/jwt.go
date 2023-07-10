package middlewares

import (
	"errors"
	"project/setting"
	"project/utils"
	"project/utils/app"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			app.ResponseErrorWithCodeMsg(c, app.CodeNoLogin, "未登陆")
			c.Abort()
			return
		}
		j := utils.NewJWT()

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				app.ResponseErrorWithCodeMsg(c, app.CodeLoginExpire, "授权过期")
				c.Abort()
				return
			}
			app.ResponseError(c, app.CodeLoginExpire)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := time.ParseDuration(setting.Conf.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
