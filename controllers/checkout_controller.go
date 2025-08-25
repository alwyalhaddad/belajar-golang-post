package controllers

import (
	"log"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/middleware"
	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Checkout(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := middleware.GetAuthenticatedUserID(c)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, "Unauthorized", "User not authenticated")
			return
		}

		// Find cart item and preload the item
		var cart models.Cart
		if err := db.Preload("CartItems").Where("user_id", userID).First(&cart).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Cart not found", "Cart item not found or not yours")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to retrive cart", err.Error())
			return
		}

		// Validate if cart is empty
		if len(cart.CartItems) == 0 {
			responses.Error(c, http.StatusBadRequest, "Cart is empty", "There are no items in cart to check out")
			return
		}

		// Start database transaction for ensure atomic operation
		tx := db.Begin()
		if tx.Error != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to start transaction", tx.Error.Error())
			return
		}
		defer tx.Rollback()

		// Calculate total amount
		var totalAmount float64
		for _, item := range cart.CartItems {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				log.Printf("Could not find product %d for checkout: %v", item.ProductID, err)
				responses.Error(c, http.StatusInternalServerError, "Failed to retrive product details", err.Error())
				return
			}
			totalAmount += product.Price * float64(item.Quantity)
		}
	}
}
