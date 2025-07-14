package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/alwyalhaddad/belajar-golang-post/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Bind request body JSON to struct LoginUserRequest
		var loginRequest models.LoginUserRequest

		err := c.ShouldBindBodyWithJSON(&loginRequest)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid Request Body", err.Error())
			return
		}

		// Find users in DB based on email
		var user models.User

		if err = db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				responses.Error(c, http.StatusUnauthorized, "Login Failed!", "Invalid email or password")
			} else {
				log.Printf("Database error finding user %s: %v", loginRequest.Email, err)
				responses.Error(c, http.StatusInternalServerError, "Login Failed!", "Internal server error")
			}
			return
		}

		// Verify Password
		if !user.CheckPasswordHash(loginRequest.Password) {
			responses.Error(c, http.StatusUnauthorized, "Login Failed!", "Invalid email or password")
			return
		}

		// If authentication success, create new session token
		sessionToken, err := utils.GenerateSessionToken(32)
		if err != nil {
			log.Printf("Failed to generate session token: %v", err)
			responses.Error(c, http.StatusInternalServerError, "Login Failed!", "Could not generate session token")
			return
		}

		// Make CSRF Token
		csrfToken, err := utils.GenerateSessionToken(16)
		if err != nil {
			log.Printf("Failed to generate CSRF token: %v", err)
			responses.Error(c, http.StatusInternalServerError, "Login Failed!", "Could not generate CSRF token")
			return
		}

		// Set session expiration time (for example, 24 hours from now)
		expiresAt := time.Now().Add(24 * time.Hour)

		// Create New Session
		newSession := models.Session{
			SessionToken: sessionToken,
			UserID:       uint(user.UserID),
			ExpiresAt:    expiresAt,
		}

		if err := db.Create(&newSession).Error; err != nil {
			log.Printf("Failed to save session to database for user %d: %v", user.UserID, err)
			responses.Error(c, http.StatusInternalServerError, "Login Failed!", "Could not create session.")
			return
		}

		//Set session token in cookie HTTP-Only
		c.SetCookie("session_token", sessionToken, int(expiresAt.Unix()-time.Now().Unix()), "/", "localhost", false, true)

		//Set email in cookie.
		c.SetCookie("email", user.Email, int(expiresAt.Unix()-time.Now().Unix()), "/", "localhost", false, false)

		responses.Success(c, http.StatusOK, "Login Success!", gin.H{
			"message":      "Login Successful",
			"email":        user.Email,
			"sessionToken": sessionToken,
			"csrfToken":    csrfToken,
			"expiresAt":    expiresAt.Format(time.RFC3339),
		})
	}
}
