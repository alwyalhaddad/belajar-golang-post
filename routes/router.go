package routes

import (
	"github.com/gin-gonic/gin"
)

func MainRoutes(router *gin.Engine) {
	mainGroup := router.Group("")
	{
		mainGroup.POST("/register")
		mainGroup.POST("/login")
		mainGroup.POST("/protected")
		mainGroup.POST("/logout")
	}
}
