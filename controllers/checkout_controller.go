package controllers

import (
	"log"
	"net/http"
	"time"

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

		// Find cart item
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

		// Make new order record
		newOrder := models.Order{
			UserID:      userID,
			TotalAmount: totalAmount,
			Status:      "Pending", // First Status
			OrderDate:   time.Now(),
		}
		if err := tx.Create(&newOrder).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to create order", err.Error())
			return
		}

		// Move item from cart to order item
		for _, cartItem := range cart.CartItems {
			orderItem := models.OrderItem{
				OrderID:   newOrder.ID,
				ProductID: cartItem.ProductID,
				Quantity:  int64(cartItem.Quantity),
				Price:     cartItem.PriceAtAddToCart,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				responses.Error(c, http.StatusInternalServerError, "Failed to create order item", err.Error())
				return
			}
		}

		// Remove all item from cart after successful moved
		if err := tx.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to clear cart", err.Error())
			return
		}

		// Commit Transaction if all operation is success
		if err := tx.Commit().Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to commit transaction", err.Error())
			return
		}

		// Give success response with detail new order
		responses.Success(c, http.StatusOK, "Checkout Success", gin.H{
			"order_id":     newOrder.ID,
			"user_id":      newOrder.UserID,
			"total_amount": newOrder.TotalAmount,
			"message":      "Your order has been successfully processed",
		})
	}
}
