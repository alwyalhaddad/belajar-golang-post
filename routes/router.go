package routes

import (
	"github.com/alwyalhaddad/belajar-golang-post/controllers"
	"github.com/alwyalhaddad/belajar-golang-post/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MainRoutes(router *gin.Engine, db *gorm.DB) {
	authMiddleware := middleware.AuthMiddleware(db)

	mainGroup := router.Group("")
	{
		mainGroup.POST("/register", controllers.Register(db))
		mainGroup.POST("/login", controllers.Login(db))
		mainGroup.POST("/logout", controllers.Logout(db))
		mainGroup.POST("/forgotpassword", controllers.ForgotPassword(db))
		mainGroup.POST("/changepassword", controllers.ChangePassword(db), authMiddleware)
	}

	productGroup := router.Group("")
	{
		productGroup.POST("/products", controllers.CreateProduct(db))
		productGroup.GET("/products", controllers.GetAllProduct(db))
		productGroup.GET("/products/:id", controllers.GetProductById(db))
		productGroup.PUT("/products/:id")
	}

	cartGroup := router.Group("cart")
	cartGroup.Use(authMiddleware)
	{
		cartGroup.POST("/items", controllers.AddItemToCart(db))
		cartGroup.GET("/items", controllers.GetCart(db))
		cartGroup.PUT("/items/:id", controllers.UpdateCartItemQuantity(db))
		cartGroup.DELETE("items/:id", controllers.RemoveCartItem(db))
	}
}
