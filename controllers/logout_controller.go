package controllers

import (
	"log"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Logout(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie
		sessionToken, err := c.Cookie("session_token")
		if err != nil {
			responses.Error(c, http.StatusUnauthorized, "Logout Failed!", "No active session")
			return
		}

		// Remove session_token from database
		result := db.Where("session_token = ?", sessionToken).Delete(&models.Session{})
		if result.Error != nil {
			log.Printf("Error deleting session %s from DB: %v", sessionToken, result.Error)
			responses.Error(c, http.StatusInternalServerError, "Logout Failed!", "Could not invalidate session in database")
			return
		}

		// Check if got any record deleted
		if result.RowsAffected == 0 {
			log.Printf("Warning: Session token %s not found in DB but received from client", sessionToken)
		}

		c.SetCookie("session_token", "", -1, "/", "localhost", false, true)
		c.SetCookie("email", "", -1, "/", "localhost", false, true)

		responses.Success(c, http.StatusOK, "Logout Success!", nil)
	}
}
