package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID           uint      `gorm:"column:id;primaryKey" json:"id"`
	SessionToken string    `gorm:"uniqueIndex;size:255;not null" json:"session_token" secure:"true" httpOnly:"true"`
	UserID       int64     `gorm:"not null" json:"user_id"`
	User         User      `gorm:"foreignKey:UserID" json:"user"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (s *Session) TableName() string {
	return "sessions"
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	if s.SessionToken == "" || s.UserID == 0 || s.ExpiresAt.IsZero() {
		return errors.New("SessionToken, UserID, and ExpiresAt cannot be empty in session")
	}
	return nil
}
