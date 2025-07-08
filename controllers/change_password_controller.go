package controllers

import (
	"log"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ChangePassword(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind request body JSON to struct ChangePassword
		var ChangePasswordRequest models.ChangePasswordRequest

		if err := c.ShouldBindBodyWithJSON(&ChangePasswordRequest); err != nil {
			responses.Error(c, http.StatusBadRequest, "Change Password Failed!", err.Error())
			return
		}

		// Validation new password and confirm
		if ChangePasswordRequest.NewPassword != ChangePasswordRequest.ConfirmNewPassword {
			responses.Error(c, http.StatusBadRequest, "Change Password Failed!", "New password and confirmation do not match.")
			return
		}

		// Checking if new password do not same with the old password
		if ChangePasswordRequest.NewPassword == ChangePasswordRequest.OldPassword {
			responses.Error(c, http.StatusBadRequest, "Change Password Failed!", "New password cannot be the same as old password.")
			return
		}

		// Make sure AuthMiddleware stores user_id in c.Get("user_id")
		userID, exists := c.Get("user_id")
		if !exists {
			log.Println("Error: UserID not found in context for change password")
			responses.Error(c, http.StatusInternalServerError, "Change Password Failed!", "Authentication context missing.")
			return
		}

		id := userID.(uint) // Convertion userID from interface{} to uint

		// Get user data from database based on user_id
		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("User with ID %d not found during change password (should not happen if authenticated)", id)
				responses.Error(c, http.StatusNotFound, "Change Password Failed!", "User not found.")
			} else {
				log.Printf("Database error retrieving %d for change password: %v", id, err)
				responses.Error(c, http.StatusInternalServerError, "Change Password Failed", "Internal server error.")
			}
			return
		}

		// Verification old password
		if !user.CheckPasswordHash(ChangePasswordRequest.OldPassword) {
			responses.Error(c, http.StatusUnauthorized, "Change Password Failed", "Incorrect old password")
			return
		}

		// New hash password
		if err := user.HashPassword(ChangePasswordRequest.NewPassword); err != nil {
			log.Printf("Error hashing new password for user %d: %v", id, err)
			responses.Error(c, http.StatusInternalServerError, "Change Password Failed!", "Could not process new password.")
			return
		}

		// Save new hash password to database
		if err := db.Model(&user).Update("PasswordHash", user.PasswordHash).Error; err != nil {
			log.Printf("Error updating password for user %d: %v", id, err)
			responses.Error(c, http.StatusInternalServerError, "Change Password Failed!", "Could not update password in database.")
			return
		}

		// Give success response
		responses.Success(c, http.StatusOK, "Change Password Success!", gin.H{
			"message": "Password changed successfuly!",
		})
	}
}
