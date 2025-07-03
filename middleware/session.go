package middleware

import (
	"errors"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/gin-gonic/gin"
)

// Dummy
var users = map[string]models.Login{
	"admin": {SessionToken: "abc123def456", CSRFToken: "token_admin"},
	"jhon":  {SessionToken: "xyz098uvw765", CSRFToken: "token_jhon"},
}

var authError = errors.New("Unauthorized")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		if username == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": authError.Error()})
			return
		}
		user, ok := users[username]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": authError.Error()})
			return
		}
		st, err := c.Cookie("session_token")
		if err != nil || st == "" || st != user.SessionToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": authError.Error()})
			return
		}
		csrf := c.GetHeader("X-CSRF-TOKEN")
		if csrf == "" || csrf != user.CSRFToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": authError.Error()})
			return
		}
		c.Next()
	}
}
