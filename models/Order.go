package models

import (
	"time"
)

type Order struct {
	ID             int64         `gorm:"column:id;primary_key" json:"id"`
	TotalAmount    float64       `gorm:"type:decimal(10,2)" json:"total_amount"`
	TotalPaid      float64       `gorm:"type:decimal(10,2)" json:"total_paid"`
	ChangeDue      float64       `gorm:"type:decimal(10,2)" json:"change_due"`
	PaymentMethod  PaymentMethod `gorm:"embedded"`
	Status         string        `gorm:"not null" json:"status"`
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

func (o *Order) TableName() string {
	return "checkouts"
}

type OrderItem struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`
	OrderID   int64     `gorm:"not null" json:"order_id"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"order"`
	ProductID int64     `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int64     `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price_per_unit"`
	Subtotal  float64   `gorm:"type:decimal(10,2);not null" json:"sub_total"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (oi *OrderItem) TableName() string {
	return "checkout_items"
}

type PaymentMethod struct {
	Cash string `json:"cash"`
	QRIS string `json:"qris"`
	Card string `json:"card"`
}

// type Status struct {
// 	Completed string `json:"completed"`
// 	Pending   string `json:"pending"`
// 	Cancelled string `json:"cancelled"`
// 	Returned  string `json:"returned"`
// }
