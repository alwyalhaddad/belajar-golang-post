package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID                 int64     `gorm:"column:user_id;primary_key" json:"userid"`
	Username               string    `gorm:"column:username;unique" json:"username" binding:"required"`
	PasswordHash           string    `gorm:"column:password_hash" json:"-"` // Don't marshal password into JSON response
	Email                  string    `gorm:"column:email;unique" json:"email" binding:"required"`
	Role                   string    `gorm:"column:role" json:"role"`
	CreatedAt              time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
	PasswordResetToken     string    `gorm:"size:255;index" json:"-"`
	PasswordResetExpiresAt time.Time `json:"-"`
}

func (u *User) tableName() string {
	return "users"
}

// HashPassword encrypts password before save to DB
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}

// CheckPassword verifies the given password a stored at hash
func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// Implement BeforeCreate for HashingPassword when create new user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Email == "" || u.PasswordHash == "" {
		return errors.New("email and password cannot be empty")
	}
	return nil
}
