package controllers

import (
	"net/http"
	"strconv"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid Product ID", "Product id must be number")
			return
		}
		var existingProduct models.Product
		if err := db.First(&existingProduct, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Product not found", "Product with that id was not found")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to retive product", err.Error())
		}
		var updateProductData models.Product
		if err := c.ShouldBindBodyWithJSON(&updateProductData); err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		}
		// Update allowed fields (avoid updating ID, CreatedAt, UpdatedAt directly)
		existingProduct.Name = updateProductData.Category.Name
		existingProduct.Description = updateProductData.Description
		existingProduct.Price = updateProductData.Price
		existingProduct.CostPrice = updateProductData.CostPrice
		existingProduct.StockQuantity = updateProductData.StockQuantity
		existingProduct.IsActive = updateProductData.IsActive
		existingProduct.CategoryID = updateProductData.CategoryID
		existingProduct.SupplierID = updateProductData.SupplierID

		// Save changes to the database
		if err := db.Save(&existingProduct).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to update product", err.Error())
			return
		}
		responses.Success(c, http.StatusOK, "Update product success", gin.H{
			"message": "Update product successfuly",
		})
	}
}
