package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// For manage shopping cart operations
type CartController struct {
	DB *gorm.DB
}

// For make new instance from CartController
func NewCartController(db *gorm.DB) *CartController {
	return &CartController{DB: db}
}

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
