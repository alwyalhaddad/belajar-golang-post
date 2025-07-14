package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var authError = errors.New("Unauthorized")

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get session token from cookie
		sessionToken, err := c.Cookie("session_token")
		if err != nil || sessionToken == "" {
			log.Printf("Auth Failed: No session token in cookie or error: %v", err)
			responses.Error(c, http.StatusUnauthorized, authError.Error(), "Session token missing or invalid.")
			c.Abort() // Important: stop the handler chain if authentication fails
			return
		}

		// Get CSRF token from header
		csrfToken := c.GetHeader("X-CSRF-TOKEN")
		if csrfToken == "" {
			log.Printf("Auth Failed: CSRF token missing in header.")
			responses.Error(c, http.StatusUnauthorized, authError.Error(), "CSRF token missing.")
			c.Abort()
			return
		}

		// Find Session at database based on sessionToken
		var session models.Session
		err = db.Where("session_token = ?", sessionToken).First(&session).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("Auth Failed: Session token '%s' not found in DB.", sessionToken)
				responses.Error(c, http.StatusUnauthorized, authError.Error(), "Invalid session token.")
			} else {
				log.Printf("Auth Failed: Database error retrieving session: %v", err)
				responses.Error(c, http.StatusInternalServerError, authError.Error(), "Internal server error during session validation.")
			}
			c.Abort()
			return
		}

		// Validate Session expiration time
		if session.IsExpired() {
			log.Printf("Auth Failed: Session token '%s' is expired.", sessionToken)
			// Remove expired session token from database
			go func() {
				if err := db.Delete(&session).Error; err != nil {
					log.Printf("Error deleting expired session %d: %v", session.ID, err)
				}
			}()
			responses.Error(c, http.StatusUnauthorized, authError.Error(), "Session expired. Please log in again.")
			c.Abort()
			return
		}

		c.Set("user_id", session.UserID)
		log.Printf("AuthMiddleware: User ID %d set in context for session %s.", session.UserID, sessionToken)

		c.Next()
	}
}
