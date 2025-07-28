package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID                     int64     `gorm:"column:id;primaryKey" json:"id"`
	Username               string    `gorm:"column:username;not null;size:100;unique" json:"username" binding:"required"`
	PasswordHash           string    `gorm:"column:password_hash" json:"-"` // Don't marshal password into JSON response
	Email                  string    `gorm:"column:email;unique" json:"email" binding:"required"`
	Role                   string    `gorm:"column:role" json:"role"`
	CreatedAt              time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	PasswordResetToken     string    `gorm:"column:password_reset_token;size:255;index;default:null" json:"-"`
	PasswordResetExpiresAt time.Time `gorm:"column:password_reset_expires_at;default:null" json:"-"`
}

func (u *User) TableName() string {
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
