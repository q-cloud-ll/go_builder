package router

import (
	"net/http"
	"project/middlewares"
	"project/setting"

	api "project/api/v1"
	"project/logger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter() *gin.Engine {
	mode := setting.Conf.Mode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 如果有跨域问题，请打开下一行
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middlewares.Cors(), middlewares.Jaeger())
	r.Use(sessions.Sessions("mysession", store))

	// 自定义路由组
	v1 := r.Group("/api/v1")

	// ---------------- 不使用jwt鉴权接口路由 ---------------
	{
		v1.POST("signup", api.UserRegisterHandler)
		v1.POST("signin", api.UserLoginHandler)
	}

	// ---------------- 使用jwt鉴权接口路由 ---------------
	v1.Use(middlewares.JWTAuth())
	{

	}
	//pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
