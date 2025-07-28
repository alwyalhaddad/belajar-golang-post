package models

import "time"

type Category struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null;size:100;unique" json:"name"`
	Description string    `gorm:"size:500" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (c *Category) TableName() string {
	return "categories"
}

type Supplier struct {
	ID            int64     `gorm:"column:id;primaryKey" json:"id"`
	Name          string    `gorm:"unique;not null;size:100" json:"name"`
	ContactPerson string    `gorm:"size:255" json:"contact"`
	PhoneNumber   int64     `gorm:"not null;size:50" json:"phone_number"`
	Email         string    `gorm:"unique;not null;size:255" json:"email"`
	Address       string    `gorm:"not null;size:500" json:"address"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (s *Supplier) TableName() string {
	return "suppliers"
}
