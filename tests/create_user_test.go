package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/alwyalhaddad/belajar-golang-post/models"
	"github.com/alwyalhaddad/belajar-golang-post/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	var Users = []models.User{
		{
			UserID:       100,
			Username:     "Admin",
			PasswordHash: "rahasia",
			Email:        "admin@example.com",
			Role:         "Admin",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			UserID:       101,
			Username:     "Jhon",
			PasswordHash: "rahasia",
			Email:        "jhon@example.com",
			Role:         "Cashier",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	for i := range Users {
		plainPassword := []byte(Users[i].PasswordHash)

		hashedPassword, err := utils.HashPassword(string(plainPassword))
		if err != nil {
			fmt.Printf("Error hashing password for user %s: %v\n",
				Users[i].Username, err)
		}
		// Update
		Users[i].PasswordHash = string(hashedPassword)
	}

	err := db.Create(&Users).Error
	assert.Nil(t, err)
}
