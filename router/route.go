package router

import (
	"net/http"
	"project/controller"
	"project/middlewares"

	"project/logger"

	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(mode string) *gin.Engine {
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
	r.Use(middlewares.Cors())

	// 自定义路由组
	v1 := r.Group("/api/v1")

	// ---------------- 不使用jwt鉴权接口路由 ---------------
	{
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("signin", controller.SignInHandler)
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
