package models

import "time"

type StockAdjustment struct {
	AdjustmentID   int64          `gorm:"column:adjustment_id;primary_key"`
	ProductID      Product        `gorm:"foreignKey:ProductId;references:product_id"`
	UserID         User           `gorm:"foreignKey:UserID;references:user_id"`
	AdjustmentType AdjustmentType `gorm:"embedded"`
	QuantityChange int64          `gorm:"column:quantity_change"`
	Reason         string         `gorm:"column:reason"`
	AdjustmentDate time.Time      `gorm:"column:adjustment_date;type:date"`
	CreatedAt      time.Time      `gorm:"column:created_at;autoCreateTime"`
}

type AdjustmentType struct {
	Add                  string `gorm:"column:add"`
	Remove               string `gorm:"column:remove"`
	Damaged              string `gorm:"column:damaged"`
	ReceivedFromSupplier string `gorm:"column:received_from_supplier"`
}
