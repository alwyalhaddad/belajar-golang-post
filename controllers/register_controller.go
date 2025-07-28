package controllers

import (
	"log"
	"net/http"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind request body JSON to struct RegisterUserRequest
		var request *models.RegisterUserRequest

		err := c.ShouldBindBodyWithJSON(&request)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
			return
		}

		// Find email is exist on database or not
		var existingUser *models.User

		if err := db.Where("username = ?", request.Username).First(&existingUser).Error; err == nil {
			responses.Error(c, http.StatusBadRequest, "Register Failed!", "Username already exist")
			return
		} else if err != gorm.ErrRecordNotFound {
			log.Printf("Database error checking username: %v", err)
			responses.Error(c, http.StatusInternalServerError, "Registration Failed!", "Internal server error")
			return
		}

		if err := db.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
			responses.Error(c, http.StatusConflict, "Registration Failed!", "Email already exists.")
			return
		} else if err != gorm.ErrRecordNotFound {
			log.Printf("Database error checking email: %v", err)
			responses.Error(c, http.StatusInternalServerError, "Registration Failed", "Internal server error")
			return
		}

		// Buat instance user baru dan hash password
		newUser := models.User{
			Username: request.Username,
			Email:    request.Email,
			Role:     request.Role,
		}
		if newUser.Role == "" {
			newUser.Role = "User"
		}

		// Hash password dari request
		if err := newUser.HashPassword(request.Password); err != nil {
			log.Printf("Error hashing password: %v", err)
			responses.Error(c, http.StatusInternalServerError, "Registration Failed", "Could not process password")
			return
		}

		// Simpan pengguna ke database
		if err := db.Create(&newUser).Error; err != nil {
			log.Printf("Error creating user in database: %v", err)
			responses.Error(c, http.StatusInternalServerError, "Registration Failed", "Could not create user")
			return
		}

		// Beri Response Success
		responses.Success(c, http.StatusCreated, "Registration Success", gin.H{
			"message":  "User registrated succesfully!",
			"user_id":  newUser.ID,
			"username": newUser.Username,
			"email":    newUser.Email,
			"role":     newUser.Role,
		})
	}
}
