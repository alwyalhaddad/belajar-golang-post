package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        int64          `gorm:"column:id;primaryKey" json:"id"`
	UserID    int64          `gorm:"not null;unique" json:"user_id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (c *Cart) TableName() string {
	return "carts"
}
