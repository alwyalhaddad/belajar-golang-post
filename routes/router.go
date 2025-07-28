package routes

import (
	"github.com/alwyalhaddad/belajar-golang-post/controllers"
	"github.com/alwyalhaddad/belajar-golang-post/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRoutes(router *gin.Engine, db *gorm.DB) {
	mainGroup := router.Group("")
	{
		mainGroup.POST("/register", controllers.Register(db))
		mainGroup.POST("/login", controllers.Login(db))
		mainGroup.POST("/logout", controllers.Logout(db))
		mainGroup.POST("/products", controllers.CreateProduct(db))
		mainGroup.GET("/products/{id}")
		mainGroup.POST("/forgotpassword", controllers.ForgotPassword(db))
		mainGroup.POST("/changepassword", middleware.AuthMiddleware(db), controllers.ChangePassword(db))
	}
}
