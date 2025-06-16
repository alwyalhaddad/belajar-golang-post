package models

import "time"

type Category struct {
	CategoryID  int64     `gorm:"column:category_id;primary_key"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at:autoCreateTime;autoUpdateTime"`
}

type Supplier struct {
	SupplierID    int64     `gorm:"column:supplier_id;primary_key"`
	Name          string    `gorm:"column:name"`
	ContactPerson string    `gorm:"column:contact_person"`
	PhoneNumber   int64     `gorm:"column:phone_number"`
	Email         string    `gorm:"column:email"`
	Address       string    `gorm:"column:address"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
