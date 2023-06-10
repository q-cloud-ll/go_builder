package router

import (
	"net/http"

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

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello,go_builder!")
		})
	}

	//pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
