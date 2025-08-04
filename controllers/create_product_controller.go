package controllers

import (
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind request body JSON to struct CreateUpdateRequest
		var request *models.CreateProductRequest

		if err := c.ShouldBindBodyWithJSON(&request); err != nil {
			responses.Error(c, http.StatusBadRequest, "Create Product Failed!", err.Error())
			return
		}

		if request.Name == "" || request.Price <= 0 || request.CostPrice <= 0 || request.StockQuantity < 0 {
			responses.Error(c, http.StatusBadRequest, "Invalid Product Data", "Value cannot be null")
			return
		}

		product := models.Product{
			Name:          request.Name,
			Description:   request.Description,
			Price:         request.Price,
			CostPrice:     request.CostPrice,
			StockQuantity: request.StockQuantity,
			IsActive:      request.IsActive,
			CategoryID:    request.CategoryID,
			SupplierID:    request.SupplierID,
		}

		// Save struct to product database
		if err := db.Create(&product).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to Create Product", err.Error())
			return
		}

		responses.Success(c, http.StatusOK, "Create Product Success!", gin.H{
			"message": "Create Product Successfuly!",
			"product": product,
		})
	}
}
