package controllers

import (
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
)

// Key is the username
var users = map[string]models.Login{}

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		responses.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}
}

var foundUser *models.User
