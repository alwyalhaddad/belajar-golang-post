package controllers

import (
	"net/http"
	"strconv"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid product id", "Product must be number")
			return
		}
		var product models.Product
		if err := db.First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Product not found", "Product with that id was not found")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
		}
		// Delete product from database
		if err := db.Delete(&product).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
		}
		responses.Success(c, http.StatusOK, "Delete product success", gin.H{
			"message": "Delete product successfuly",
		})
	}
}
