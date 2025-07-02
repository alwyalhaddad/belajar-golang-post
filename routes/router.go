package routes

import (
	"github.com/alwyalhaddad/belajar-golang-post/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRoutes(router *gin.Engine, db *gorm.DB) {
	mainGroup := router.Group("")
	{
		mainGroup.POST("/register", controllers.Register(db))
		mainGroup.POST("/login", controllers.Login(db))
		mainGroup.POST("/changepassword", controllers.ChangePassword(db))
		mainGroup.POST("/logout", controllers.Logout(db))
	}
}
