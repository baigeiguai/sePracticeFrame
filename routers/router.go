package routers

import (
	"net/http"
	"sePracticeFrame/controller"
	"sePracticeFrame/settings"

	"github.com/gin-gonic/gin"
)

func SetupRouters() (r *gin.Engine) {
	if settings.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r = gin.Default()
	// r.Use(middlewares.Cores())
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	userGroup := r.Group("user")
	// userGroup.Use(middlewares.Cores())
	{
		userGroup.POST("/register", controller.UserRigersterHandler)
		userGroup.POST("/login", controller.UserLoginHandler)
		userGroup.OPTIONS("/register", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})
		userGroup.OPTIONS("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})
	}
	return
}
