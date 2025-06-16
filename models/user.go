package models

import "time"

type User struct {
	UserID       int64     `gorm:"column:user_id;primary_key;"`
	Username     string    `gorm:"column:username"`
	PasswordHash int64     `gorm:"column:password_hash"`
	Email        string    `gorm:"email"`
	Role         Role      `gorm:"embedded"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

type Role struct {
	Admin   string `gorm:"column:admin"`
	Manager string `gorm:"column:manager"`
	Cashier string `gorm:"column:cashier"`
}
