package controller

import (
	"net/http"
	"sePracticeFrame/models"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func UserRigersterHandler(c *gin.Context) {
	var user models.User= models.User{}
	c.BindJSON(&user)
	registerState := models.RegisterUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  registerState,
		"message": registerState.String(),
	})
}
func UserLoginHandler(c *gin.Context) {
	var user models.User=models.User{}
	c.BindJSON((&user))
	loginState := models.LoginUser(&user)
	c.JSON(http.StatusOK,gin.H{
		"status":loginState,
		"message": loginState.String(),
	})
}
