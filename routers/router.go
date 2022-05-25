package routers

import (
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
	{
		userGroup.POST("/register", controller.UserRigersterHandler)
		userGroup.POST("/login", controller.UserLoginHandler)
	}
	return
}
