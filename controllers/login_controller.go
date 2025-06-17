package controllers

import (
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err == nil {
		responses.Success(c, http.StatusCreated, "Login Success!", err)
	} else {
		responses.Error(c, http.StatusBadRequest, "Login Failed!", err.Error())
		return
	}
}
