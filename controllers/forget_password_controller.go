package controllers

import (
	"log"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ForgotPassword(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind request body JSON to struct ForgotPassword
		var ForgotPasswordRequest models.ForgotPasswordRequest

		if err := c.ShouldBindBodyWithJSON(&ForgotPasswordRequest); err != nil {
			responses.Error(c, http.StatusBadRequest, "Forgot Password Failed!", err.Error())
			return
		}

		// Find user by email
		var user models.User

		if err := db.Where("email = ?", ForgotPasswordRequest.Email).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("Forgot password: Email %s not found, but sending success responses for security.", ForgotPasswordRequest.Email)
				responses.Success(c, http.StatusOK, "Forgot Password Request Received!", "If your email is registered, you will receive a password link sortly.")
				return
			} else {
				log.Printf("Database error finding user by email %s: %v", ForgotPasswordRequest.Email)
				responses.Error(c, http.StatusInternalServerError, "Forgot Password Failed!", "Internal server error")
				return
			}
		}

		// Use a longer, more secure GenerateToken
		resetToken := models.GenerateSessionToken(32)
	}
}
