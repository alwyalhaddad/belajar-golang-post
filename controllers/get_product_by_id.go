package controllers

import (
	"net/http"
	"strconv"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid Product ID", "id should be number")
			return
		}

		var request models.Product
		if err := db.Preload("Category").Preload("Supplier").First(&request, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusNotFound, "Product not found", "Product with that ID was not found")
				return
			}
			responses.Error(c, http.StatusInternalServerError, "Failed to retrive product", err.Error())
			return
		}
		responses.Success(c, http.StatusOK, "Get product success", gin.H{
			"message": "Get product by id successduly",
		})
	}
}
