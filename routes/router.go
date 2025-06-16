package routes

import (
	"github.com/alwyalhaddad/belajar-golang-post/controllers"
	"github.com/gin-gonic/gin"
)

func MainRoutes(router *gin.Engine) {
	mainGroup := router.Group("")
	{
		mainGroup.POST("/login", controllers.Login)
	}
}
