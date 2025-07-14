package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/alwyalhaddad/belajar-golang-post/utils"
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
				log.Printf("Database error finding user by email %s: %v", ForgotPasswordRequest.Email, err)
				responses.Error(c, http.StatusInternalServerError, "Forgot Password Failed!", "Internal server error")
				return
			}
		}

		// Use a longer, more secure GenerateToken
		resetToken, _ := utils.GenerateSessionToken(32)

		// Set time token expires to database for user
		expiresAt := time.Now().Add(1 * time.Hour)

		// Save token and time expires to database
		if err := db.Model(&user).Updates(map[string]interface{}{
			"PasswordResetToken":     resetToken,
			"PasswordResetExpiresAt": expiresAt,
		}).Error; err != nil {
			log.Printf("Error saving password reset token for user %d: %v", user.UserID, err)
			responses.Error(c, http.StatusInternalServerError, "Forgot Password Failed!", "Could not generate token.")
			return
		}

		// Send email reset password to user
		resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", resetToken)
		emailBody := fmt.Sprintf("Dear %s,\n\n", user.Username) +
			fmt.Sprintf("You have requested to reset your password. Please click on the link below to reset your password:\n%s\n\n", resetLink) +
			"This link will expire in 1 hour. If you did not request this, please ignore this email.\n\n" +
			"Thanks,\n\nBlockchainStore"

		if err := utils.SendPasswordResetEmail(user.Email, "Password Reset Request", emailBody); err != nil {
			log.Printf("Failed to send reset email %s: %v", user.Email, err)
		}

		// Response success
		responses.Success(c, http.StatusOK, "Forgot Password Received!", "If your email is registered, you will receive a password reset link shortly.")
	}
}
