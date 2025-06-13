package database

import "time"

type Customer struct {
	CustomerID   int64     `gorm:"column:customer_id;primary_key"`
	Name         Name      `gorm:"embedded"`
	PhoneNumber  int64     `gorm:"column:phone_number"`
	Email        string    `gorm:"column:email"`
	Address      string    `gorm:"column:address"`
	LoyaltyPoint int64     `gorm:"column:loyalty_point"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}
