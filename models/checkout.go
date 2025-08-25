package models

import (
	"time"
)

type Checkout struct {
	ID             int64         `gorm:"column:id;primary_key" json:"id"`
	TotalAmount    float64       `gorm:"type:decimal(10,2)" json:"total_amount"`
	TotalPaid      float64       `gorm:"type:decimal(10,2)" json:"total_paid"`
	ChangeDue      float64       `gorm:"type:decimal(10,2)" json:"change_due"`
	PaymentMethod  PaymentMethod `gorm:"embedded"`
	Status         Status        `gorm:"embedded"`
	UserID         int64         `gorm:"not null" json:"user_id"`
	User           User          `gorm:"foreignKey:UserID" json:"user"`
	CustomerID     *uint         `gorm:"not null" json:"customer_id"`
	Customer       Customer      `gorm:"foreignKey:CustomerID" json:"customer"`
	DiscountAmount float64       `gorm:"type:decimal(10,2)" json:"discount_amount"`
	TaxAmount      float64       `gorm:"type:decimal(10,2)" json:"tax_amount"`
	OrderDate      time.Time     `gorm:"type:date;index;not null" json:"order_date"`
	CreatedAt      time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

func (o *Checkout) TableName() string {
	return "checkouts"
}

type CheckoutItem struct {
	ID           int64     `gorm:"column:id;primaryKey" json:"id"`
	CheckoutID   int64     `gorm:"not null" json:"checkout_id"`
	Checkout     Checkout  `gorm:"foreignKey:CheckoutID" json:"checkout"`
	ProductID    int64     `gorm:"not null" json:"product_id"`
	Product      Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity     int64     `gorm:"not null" json:"quantity"`
	PricePerUnit float64   `gorm:"type:decimal(10,2);not null" json:"price_per_unit"`
	Subtotal     float64   `gorm:"type:decimal(10,2);not null" json:"sub_total"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func (oi *CheckoutItem) TableName() string {
	return "checkout_items"
}

type PaymentMethod struct {
	Cash string `json:"cash"`
	QRIS string `json:"qris"`
	Card string `json:"card"`
}

type Status struct {
	Completed string `json:"completed"`
	Pending   string `json:"pending"`
	Cancelled string `json:"cancelled"`
	Returned  string `json:"returned"`
}
