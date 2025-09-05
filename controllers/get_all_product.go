package controllers

import (
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request []models.Product
		if err := db.Preload("Category").Preload("Supplier").Find(&request).Error; err != nil {
			responses.Error(c, http.StatusInternalServerError, "Failed to retrive product", err.Error())
			return
		}

		responses.Success(c, http.StatusOK, "Get product success", gin.H{
			"message":  "Get all product successfuly",
			"products": request,
		})
	}
}
