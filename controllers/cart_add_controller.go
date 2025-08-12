package controllers

import (
	"log"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddItemToCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind request body JSON to struct
		var request models.CartRequest
		err := c.ShouldBindBodyWithJSON(&request)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
			return
		}

		if request.ProductID == 0 || request.Quantity <= 0 {
			responses.Error(c, http.StatusBadRequest, "Invalid Input", "Product ID and Quantity must be valid")
			return
		}

		userID, err := GetAuthenticatedUserID(c)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, "Unauthorized", "User not authenticated")
			return
		}

		// Find or create cart for new user
		var cart models.Cart
		if err := db.Where("user_id = ?", userID).First(&cart).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				cart := models.Cart{UserID: userID}
				if err := db.Create(&cart).Error; err != nil {
					responses.Error(c, http.StatusInternalServerError, "Failed to create cart", err.Error())
					return
				}
			} else {
				responses.Error(c, http.StatusInternalServerError, "Failed to retrive cart", err.Error())
				return
			}
		}
		// Find Product
		var product models.Product
		if err := db.First(&product, request.ProductID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Product Not Found", "Product not found")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to retrive product", err.Error())
			return
		}

		// Serach for existing cart item for this product
		var cartItem models.CartItem
		found := true
		if err := db.Where("cart_id = ? AND product_id = ?", cart.ID, request.ProductID).First(&cartItem).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				found = false
			} else {
				responses.Error(c, http.StatusInternalServerError, "Failed to check cart item", err.Error())
				return
			}
		}
		if found {
			// Update quantity if item is exist
			cartItem.Quantity += int(request.Quantity)
			if err := db.Save(&cartItem).Error; err != nil {
				responses.Error(c, http.StatusInternalServerError, "Failed to update cart item quantity", err.Error())
				return
			}
		} else {
			cartItem = models.CartItem{
				CartID:           cart.ID,
				ProductID:        product.ID,
				Quantity:         int(request.Quantity),
				PriceAtAddToCart: product.Price,
			}
			if err := db.Create(&cartItem).Error; err != nil {
				responses.Error(c, http.StatusInternalServerError, "Failed to add item to cart", err.Error())
				return
			}
		}
		// Preload cart item with detail product for response
		if err := db.Preload("Product").First(&cartItem, cartItem.ID).Error; err != nil {
			log.Printf("Warning: Could not preload for cart item %d: %v", cartItem.ID, err)
		}
		responses.Success(c, http.StatusOK, "Add item to cart success!", gin.H{
			"message": "Add item to cart successfuly",
		})
	}
}
