package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SessionToken string    `gorm:"uniqueIndex;size:255;not null" json:"session_token" secure:"true" httpOnly:"true"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User         User      `gorm:"foreignKey:UserID;references:user_id"`
}

func (s *Session) tableName() string {
	return "sessions"
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	if s.SessionToken == "" {
		return errors.New("SessionToken cannot be empty")
	}
	if s.UserID == 0 {
		return errors.New("userID cannot be 0")
	}
	if s.ExpiresAt.Before(time.Now()) {
		return errors.New("ExpiresAt must be in the future")
	}
	return
}
