package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAuthenticatedUserID(c *gin.Context) (int64, error) {
	UserID, exist := c.Get("UserID")
	if !exist {
		return 0, gorm.ErrRecordNotFound
	}

	// Convert value to int64
	id, ok := UserID.(int64)
	if !ok {
		return 0, gorm.ErrInvalidData
	}

	return id, nil
}

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

// Handle retrieving content to users cart
func GetCart(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := GetAuthenticatedUserID(c)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, "Unauthorized", "User not authenticated")
			return
		}
		var cart models.Cart
		if err := db.Preload("CartItems.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Cart not found", "The shopping cart is empty or has not been created")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to retrive cart", err.Error())
			return
		}
		responses.Success(c, http.StatusOK, "Get Cart Successfuly", gin.H{
			"message": "Get Cart Successfuly",
		})
	}
}

func UpdateCartItemQuantity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cartItemIDStr := c.Param("id")
		cartItemID, err := strconv.ParseUint(cartItemIDStr, 10, 64)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid request body", "Item id should be number")
			return
		}
		// Bind request body JSON to struct
		var request models.CartRequest

		if err := c.ShouldBindBodyWithJSON(&request); err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
			return
		}

		if request.Quantity < 0 {
			responses.Error(c, http.StatusBadRequest, "Invalid quantity", "Quantitas cant be negative")
			return
		}

		userID, err := GetAuthenticatedUserID(c)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, "Unauthorized", "User not authenticated")
			return
		}
		// Make sure cart item belong to correct user & correct shopping cart
		var cartItem models.CartItem

		err = db.Joins("JOIN cart ON cart.id = cart_items.cart_id").
			Where("cart_item = ? AND carts.user_id = ?", cartItemID, userID).First(&cartItem).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Cart item not found", "Cart item not found or not yours")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to retrive cart item", err.Error())
			return
		}

		if request.Quantity == 0 {
			if err := db.Delete(&cartItem).Error; err != nil {
				responses.Error(c, http.StatusInternalServerError, "Failed to remove item from cart", err.Error())
				return
			}
			responses.Error(c, http.StatusNoContent, "Status no content", "Status no content")
			return
		}
		cartItem.Quantity = int(request.Quantity)
		if err := db.Save(&cartItem).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to update cart item", err.Error())
			return
		}
		if err := db.Preload("Product").First(&cartItem, cartItemID).Error; err != nil {
			log.Printf("Warning: Could not preload product for cart item %d: %v", cartItem.ID, err)
		}
		responses.Success(c, http.StatusOK, "Update cart item successfuly", gin.H{
			"message": "Update cart item successfuly",
		})
	}
}

func RemoveCartItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cartItemIDStr := c.Param("id")
		cartItemID, err := strconv.ParseUint(cartItemIDStr, 10, 64)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid cart item id", "Cart item id must be number")
			return
		}

		userID, err := GetAuthenticatedUserID(c)
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, "Unauthorized", "User not authenticated")
			return
		}

		// Make sure cart item is belong to correct owner & correct shopping cart
		var cartItem models.CartItem

		err = db.Joins("JOIN carts on carts.id = cart_items.cart_id").
			Where("cart_items.id = ? AND carts.user_id = ?", cartItemID, userID).
			First(&cartItem).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Cart item not found", "Cart item not found or not yours ")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to retrive cart", err.Error())
			return
		}

		if err := db.Delete(&cartItem).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to delete cart item", err.Error())
			return
		}

		c.Status(http.StatusNoContent) // 204 no content for success deleting cart item
	}
}
