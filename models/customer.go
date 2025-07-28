package models

import "time"

type Customer struct {
	ID           int64     `gorm:"column:id;primaryKey" json:"id"`
	Name         Name      `gorm:"column:name;not null;size:100;unique" json:"name"`
	PhoneNumber  int64     `gorm:"not null;size:50" json:"phone_number"`
	Email        string    `gorm:"unique;not null;size:500" json:"email"`
	Address      string    `gorm:"not null;size500" json:"address"`
	LoyaltyPoint int64     `gorm:"column:loyalty_point"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (c *Customer) TableName() string {
	return "customers"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}
