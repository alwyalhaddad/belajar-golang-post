package database

import (
	"time"
)

type Order struct {
	OrderID        int64         `gorm:"column:order_id;primary_key"`
	TotalAmount    int64         `gorm:"column:total_amount"`
	TotalPaid      int64         `gorm:"column:total_paid"`
	ChangeDue      int64         `gorm:"column:change_due"`
	PaymentMethod  PaymentMethod `gorm:"embedded"`
	Status         Status        `gorm:"embedded"`
	UserID         User          `gorm:"foreignKey:UserID;references:user_id"`
	CustomerID     *int64        // FK (automatis jika nil)
	Customer       *Customer
	DiscountAmount int64     `gorm:"column:discount_amount"`
	TaxAmount      int64     `gorm:"column:tax_amount"`
	OrderDate      time.Time `gorm:"column:order_date;type:date;index"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

type PaymentMethod struct {
	Cash string `gorm:"column:cash"`
	QRIS string `gorm:"column:qris"`
	Card string `gorm:"column:card"`
}

type Status struct {
	Completed string `gorm:"column:completed"`
	Pending   string `gorm:"column:pending"`
	Cancelled string `gorm:"column:cancelled"`
	Returned  string `gorm:"column:returned"`
}

type OrderItem struct {
	OrderItemID  int64     `gorm:"column:order_item_id;primary_key"`
	OrderID      Order     `gorm:"foreignKey:OrderID;references:order_id"`
	ProductID    Product   `gorm:"foreignKey:ProductID;references:product_id"`
	Quantity     int64     `gorm:"column:quantity"`
	PricePerUnit int64     `gorm:"column:price_per_unit"`
	Subtotal     int64     `gorm:"column:subtotal"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}
