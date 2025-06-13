package database

import "time"

type Return struct {
	ReturnID     int64     `gorm:"column:return_id;primary_key"`
	OrderID      Order     `gorm:"foreignKey:OrderID;references:order_id"`
	ReturnAmount int64     `gorm:"column:return_amount"`
	Reason       string    `gorm:"column:reason"`
	UserID       User      `gorm:"foreignKey"`
	ReturnDate   time.Time `gorm:"column:return_date;type:date"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

type ReturnItem struct {
	ReturnItemID int64     `gorm:"column:return_item_id;primary_key"`
	ReturnID     Return    `gorm:"foreignKey:ReturnID;references:return_id"`
	ProductID    Product   `gorm:"foreignKey:ProductID;reference:product_id"`
	Quantity     int64     `gorm:"column:quantity"`
	PricePerUnit int64     `gorm:"column:price_per_unit"`
	Subtotal     int64     `gorm:"column:subtotal"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}
