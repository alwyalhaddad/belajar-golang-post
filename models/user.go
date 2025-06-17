package models

import "time"

type User struct {
	UserID       int64     `gorm:"column:user_id;primary_key" json:"userid"`
	Username     string    `gorm:"column:username" json:"username" binding:"required"`
	PasswordHash string    `gorm:"column:password_hash" json:"-"` // Don't marshal password into JSON response
	Email        string    `gorm:"email" json:"email" binding:"required"`
	Role         string    `gorm:"column:role" json:"role"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
}
