package models

import "time"

type Return struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	CheckoutID   int64     `gorm:"not null" json:"checkout_id"`
	Order        Order     `gorm:"foreignKey:CheckoutID" json:"order"`
	UserID       int64     `gorm:"not null" json:"user_id"`
	User         User      `gorm:"foreignKey:UserID" json:"user"`
	ReturnAmount int64     `gorm:"not null" json:"return_amount"`
	Reason       string    `gorm:"not null" json:"reason"`
	ReturnDate   time.Time `gorm:"type:date"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (r *Return) TableName() string {
	return "returns"
}

type ReturnItem struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	ReturnID     int64     `gorm:"not null" json:"return_id"`
	Return       Return    `gorm:"foreignKey:ReturnID" json:"return"`
	ProductID    int64     `gorm:"not null" json:"product_id"`
	Product      Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity     int64     `gorm:"not null" json:"quantity"`
	PricePerUnit float64   `gorm:"type:decimal(10,2);not null" json:"price_per_unit"`
	Subtotal     float64   `gorm:"type:decimal(10,2);not null" json:"sub_total"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (ri *ReturnItem) TableName() string {
	return "return_items"
}
